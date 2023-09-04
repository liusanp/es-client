package router

import (
	"es-client/api"

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
	r.GET("/index", api.GetIndex)
	// 配置
	r.GET("/conf/get", api.GetConfig)
	r.POST("/conf/set", api.SetConfig)
	r.POST("/conf/use", api.UseConfig)
	return r
}
