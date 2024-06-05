package controller

import (
	"es-client/commons"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	BaseController
}

// GetIndices
// @Summary 获取es索引
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/es/getIndices [get]
func (con BaseController) GetIndices(c *gin.Context) {
	indices, err := commons.GetIndices()
	if err != nil {
		con.Error(c, "获取索引列表失败", err.Error())
		return
	}

	con.Ok(c, "获取索引列表成功", indices)
}

// GetMappings
// @Summary 获取es索引mappings
// @Tags es查询
// @Accept json
// @Produce json
// @Param        index    query     string  false  "索引名称"
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/es/getMappings [get]
func (con BaseController) GetMappings(c *gin.Context) {
	index := c.Query("index")
	mappings, err := commons.GetMappings(index)

	if err != nil {
		con.Error(c, "获取索引mappings失败", err.Error())
		return
	}

	con.Ok(c, "获取索引mappings成功", mappings)
}
