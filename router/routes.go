package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/thiagoCalazans-dev/go-api-opportunities.git/docs"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.Init()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := r.Group(basePath)
	{
		v1.GET("/opportunity", handler.ShowOportunity)
		v1.POST("/opportunity", handler.CreateOportunity)
		v1.PUT("/opportunity", handler.UpdateOportunity)
		v1.DELETE("/opportunity", handler.RemoveOportunity)
		v1.GET("/opportunities", handler.ListOportunity)
	}
	//initialize swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
