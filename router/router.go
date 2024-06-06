package router

import (
	"es-client/controller"
	"es-client/public"

	docs "es-client/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 静态资源
	public.HandleStatic(&r.RouterGroup, r.NoRoute)
	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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
