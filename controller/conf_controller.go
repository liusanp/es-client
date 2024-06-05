package controller

import (
	"es-client/commons"
	"es-client/models"

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
// @Router /ec/conf/get [get]
func (con ConfController) GetConfig(c *gin.Context) {
	conf := commons.GetESConfigs()
	con.Ok(c, "获取配置成功", conf)
}

// AddConfig
// @Summary 添加es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Param        newConfig    body     models.ESConfig  false  "es配置"
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/conf/add [post]
func (con ConfController) AddConfig(c *gin.Context) {
	var newConfig models.ESConfig
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		con.Error(c, "添加配置失败", err.Error())
		return
	}
	newConfig.Selected = false
	err := commons.AddESConfig(newConfig)
	if err != "" {
		con.Error(c, "添加配置失败", err)
		return
	}
	con.Ok(c, "添加配置成功", newConfig)
}

// DeleteConfig
// @Summary 删除es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Param        name    query     string  false  "索引名称"
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/conf/del [post]
func (con ConfController) DeleteConfig(c *gin.Context) {
	name := c.Query("name")
	err := commons.DeleteESConfig(name)
	if err != "" {
		con.Error(c, "删除配置失败", err)
		return
	}
	con.Ok(c, "删除配置成功", nil)
}

// SelectConfig
// @Summary 设置es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Param        newConfig    body     models.ESConfig  false  "es配置"
// @Success 200 {string} json{"code","msg","data"}
// @Router /ec/conf/use [post]
func (con ConfController) SelectConfig(c *gin.Context) {
	var newConfig models.ESConfig
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		con.Error(c, "启用配置失败", err.Error())
		return
	}

	currentConfig, err := commons.SelectESConfig(newConfig)
	if err != nil {
		con.Error(c, "启用配置失败", err.Error())
		return
	}
	con.Ok(c, "启用配置成功", currentConfig)
}
