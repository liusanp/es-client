package controller

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

// GetIndex
// @Summary 首页
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router / [get]
func (con IndexController) GetIndex(c *gin.Context) {
	con.Ok(c, "首页", "首页")
}
