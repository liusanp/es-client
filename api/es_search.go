package api

import (
	"es-client/commons"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMapping
// @Summary 获取索引字段
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json {"code","msg","data"}
// @Router /es/getMapping [get]
func GetMapping(c *gin.Context) {
	commons.GetIndexMapping()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取字段成功",
		"data": "",
	})
}