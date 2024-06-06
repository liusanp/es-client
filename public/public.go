package public

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:embed all:dist
var Public embed.FS

var static fs.FS

func initStatic() {
	dist, err := fs.Sub(Public, "dist")
	if err != nil {
		log.Fatalf("failed to read dist dir")
	}
	static = dist
}

func initIndex() string {
	indexFile, err := static.Open("index.html")
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Fatalf("index.html not exist, you may forget to put dist of frontend to public/dist")
		}
		log.Fatalf("failed to read index.html: %v", err)
	}
	defer func() {
		_ = indexFile.Close()
	}()
	index, err := io.ReadAll(indexFile)
	if err != nil {
		log.Fatalf("failed to read dist/index.html")
	}
	return string(index)
}

func HandleStatic(r *gin.RouterGroup, noRoute func(handlers ...gin.HandlerFunc)) {
	initStatic()
	indexHtml := initIndex()
	folders := []string{"assets", "browser_upgrade"}
	r.Use(func(c *gin.Context) {
		for i := range folders {
			if strings.HasPrefix(c.Request.RequestURI, fmt.Sprintf("/%s/", folders[i])) {
				c.Header("Cache-Control", "public, max-age=15552000")
			}
		}
	})
	for i, folder := range folders {
		sub, err := fs.Sub(static, folder)
		if err != nil {
			log.Fatalf("can't find folder: %s", folder)
		}
		r.StaticFS(fmt.Sprintf("/%s/", folders[i]), http.FS(sub))
	}
	r.GET("/favicon.ico", func(c *gin.Context) {
		favicon, _ := Public.ReadFile("dist/favicon.ico")
		c.Data(http.StatusOK, "image/x-icon", favicon)
	})
	r.GET("/loading.css", func(c *gin.Context) {
		lc, _ := Public.ReadFile("dist/loading.css")
		c.Data(http.StatusOK, "text/css; charset=utf-8", lc)
	})

	noRoute(func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.Status(200)
		_, _ = c.Writer.WriteString(indexHtml)
		c.Writer.Flush()
		c.Writer.WriteHeaderNow()
	})
}
