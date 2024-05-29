package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.Init()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opportunity", handler.ShowOportunity)
		v1.POST("/opportunity", handler.CreateOportunity)
		v1.PUT("/opportunity", handler.UpdateOportunity)
		v1.DELETE("/opportunity", handler.RemoveOportunity)
		v1.GET("/opportunities", handler.ListOportunity)
	}
}
