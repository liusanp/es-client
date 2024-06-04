package service

import (
	"es-client/commons"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndices
// @Summary 获取es索引
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /indices/get [get]
func GetIndices(c *gin.Context) {
	indices, err := commons.GetIndices()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "获取索引列表失败",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取索引列表成功",
		"data": indices,
	})
}

// GetMappings
// @Summary 获取es索引mappings
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /indices/mappings [get]
func GetMappings(c *gin.Context) {
	index := c.Param("index")
	mappings, err := commons.GetMappings(index)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "获取索引mappings失败",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取索引mappings成功",
		"data": mappings,
	})
}
