package service

import (
	"es-client/commons"
	"es-client/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetConfig
// @Summary 获取es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Success 200 {string} json {"code","msg","data"}
// @Router /conf/get [get]
func GetConfig(c *gin.Context) {
	conf := commons.GetConfig("app.es.conf")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取配置成功",
		"data": conf,
	})
}

// SetConfig
// @Summary 设置es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Param esConf body []models.EsConfig true "EsConfig"
// @Success 200 {string} json {"code","msg","data"}
// @Router /conf/set [post]
func SetConfig(c *gin.Context) {
	// conf := commons.GetConfig("app.es.conf")
	esConf := []models.EsConfig{}
	if err := c.BindJSON(&esConf); err != nil {
		log.Println(err)
	}
	log.Println(esConf)
	commons.SetConfig("app.es.conf", esConf)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "设置配置成功",
		"data": nil,
	})
}

// UseConfig
// @Summary 应用es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Param esConf body models.EsConfig true "EsConfig"
// @Success 200 {string} json {"code","msg","data"}
// @Router /conf/use [post]
func UseConfig(c *gin.Context) {
	// conf := commons.GetConfig("app.es.conf")
	esConf := models.EsConfig{}
	if err := c.BindJSON(&esConf); err != nil {
		log.Println(err)
	}
	log.Println(esConf)
	commons.InitESClient(&esConf)
	commons.CheckESClient(esConf.Version)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "应用配置成功",
		"data": nil,
	})
}