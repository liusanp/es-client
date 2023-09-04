package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Summary 首页
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "请求成功",
		"data": "es-client",
	})
}
