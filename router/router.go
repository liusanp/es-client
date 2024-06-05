package router

import (
	"es-client/controller"

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
	r.GET("/", controller.IndexController{}.GetIndex)
	// 配置
	confRoute := r.Group("/ec/conf")
	{
		confRoute.GET("/get", controller.ConfController{}.GetConfig)
		confRoute.POST("/add", controller.ConfController{}.AddConfig)
		confRoute.POST("/del", controller.ConfController{}.DeleteConfig)
		confRoute.POST("/use", controller.ConfController{}.SelectConfig)
	}
	// 查询
	esRoute := r.Group("/ec/es")
	{
		esRoute.GET("/getIndices", controller.SearchController{}.GetIndices)
		esRoute.GET("/getMappings", controller.SearchController{}.GetMappings)
	}
	return r
}
