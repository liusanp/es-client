package service

import (
	"es-client/commons"
	"es-client/models"
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
	config := commons.GetESConfigs()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取配置成功",
		"data": config,
	})
}

// AddConfig
// @Summary 添加es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /conf/add [post]
func AddConfig(c *gin.Context) {
	var newConfig models.ESConfig
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "添加配置失败",
			"data": nil,
		})
		return
	}

	err := commons.AddESConfig(newConfig)
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "添加配置失败",
			"data": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "添加配置成功",
		"data": newConfig,
	})
}

// DeleteConfig
// @Summary 删除es配置
// @Tags es配置
// @Accept string
// @Produce string
// @Success 200 {string} json{"code","msg","data"}
// @Router /conf/del [post]
func DeleteConfig(c *gin.Context) {
	name := c.Param("name")
	err := commons.DeleteESConfig(name)
	if err != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "删除配置失败",
			"data": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除配置成功",
		"data": nil,
	})
}

// SelectConfig
// @Summary 设置es配置
// @Tags es配置
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","msg","data"}
// @Router /conf/set [post]
func SelectConfig(c *gin.Context) {
	var newConfig models.ESConfig
	if err := c.ShouldBindJSON(&newConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10001,
			"msg":  "设置配置失败",
			"data": err.Error(),
		})
		return
	}

	currentConfig := commons.SelectESConfig(newConfig)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "设置配置成功",
		"data": currentConfig,
	})
}
