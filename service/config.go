package service

import (
	"es-client/commons"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetConfig
// @Summary 获取es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /conf/get [get]
func GetConfig(c *gin.Context) {
	conf := commons.GetConfig("app.es.conf")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取配置成功",
		"data": conf,
	})
}
