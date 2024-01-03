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
	confRoute := r.Group("/conf")
	{
		confRoute.GET("/get", controller.ConfController{}.GetConfig)
		confRoute.POST("/set", controller.ConfController{}.SetConfig)
		confRoute.POST("/use", controller.ConfController{}.UseConfig)
	}
	// 查询
	esRoute := r.Group("/es")
	{
		esRoute.GET("/getIndices", controller.SearchController{}.GetIndices)
		esRoute.POST("/getMapping", controller.SearchController{}.GetMapping)
	}
	return r
}
