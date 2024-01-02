package controller

import (
	"es-client/commons"

	"github.com/gin-gonic/gin"
)

type SearchController struct {
	BaseController
}

// GetIndices
// @Summary 获取索引
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json {"code","msg","data"}
// @Router /es/getIndices [get]
func (con BaseController) GetIndices(c *gin.Context) {
	res := commons.GetIndices()
	con.Ok(c, "获取索引成功", res)
}

// GetMapping
// @Summary 获取索引字段
// @Tags es查询
// @Accept json
// @Produce json
// @Success 200 {string} json {"code","msg","data"}
// @Router /es/getMapping [get]
func (con BaseController) GetMapping(c *gin.Context) {
	res := commons.GetIndexMapping(c.GetString("index"))
	con.Ok(c, "获取字段成功", res)
}
