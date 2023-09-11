package api

import (
	"es-client/commons"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndices
// @Summary 获取索引
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json {"code","msg","data"}
// @Router /es/getIndices [get]
func GetIndices(c *gin.Context) {
	res := commons.GetIndices()

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取索引成功",
		"data": res,
	})
}

// GetMapping
// @Summary 获取索引字段
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json {"code","msg","data"}
// @Router /es/getMapping [get]
func GetMapping(c *gin.Context) {
	res := commons.GetIndexMapping(c.GetString("index"))

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取字段成功",
		"data": res,
	})
}