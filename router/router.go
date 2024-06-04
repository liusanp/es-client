package router

import (
	"es-client/service"

	docs "es-client/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 首页
	r.GET("/index", service.GetIndex)
	// 配置
	r.GET("/conf/get", service.GetConfig)
	r.POST("/conf/add", service.AddConfig)
	r.POST("/conf/del", service.DeleteConfig)
	r.POST("/conf/set", service.SelectConfig)
	// 查询
	r.GET("/indices/get", service.GetIndices)
	r.GET("/indices/mappings", service.GetMappings)
	return r
}
