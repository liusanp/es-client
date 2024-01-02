package controller

import (
	"es-client/commons"
	"es-client/models"
	"log"

	"github.com/gin-gonic/gin"
)

type ConfController struct {
	BaseController
}

// GetConfig
// @Summary 获取es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Success 200 {string} json {"code","msg","data"}
// @Router /conf/get [get]
func (con ConfController) GetConfig(c *gin.Context) {
	conf := commons.GetConfig("app.es.conf")
	con.Ok(c, "获取配置成功", conf)
}

// SetConfig
// @Summary 设置es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Param esConf body []models.EsConfig true "EsConfig"
// @Success 200 {string} json {"code","msg","data"}
// @Router /conf/set [post]
func (con ConfController) SetConfig(c *gin.Context) {
	// conf := commons.GetConfig("app.es.conf")
	esConf := []models.EsConfig{}
	if err := c.BindJSON(&esConf); err != nil {
		log.Println(err)
	}
	log.Println(esConf)
	commons.SetConfig("app.es.conf", esConf)
	con.Ok(c, "设置配置成功", nil)
}

// UseConfig
// @Summary 应用es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Param esConf body models.EsConfig true "EsConfig"
// @Success 200 {string} json {"code","msg","data"}
// @Router /conf/use [post]
func (con ConfController) UseConfig(c *gin.Context) {
	// conf := commons.GetConfig("app.es.conf")
	esConf := models.EsConfig{}
	if err := c.BindJSON(&esConf); err != nil {
		log.Println(err)
	}
	log.Println(esConf)
	commons.InitESClient(&esConf)
	info := commons.CheckESClient()
	con.Ok(c, "应用配置成功", info)
}
