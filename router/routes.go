package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/opportunity", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "get opportunity",
			})
		})
		v1.POST("/opportunity", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "get opportunity",
			})
		})
		v1.PUT("/opportunity", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "get opportunity",
			})
		})
		v1.DELETE("/opportunity", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "get opportunity",
			})
		})
		v1.GET("/opportunities", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "get opportunity",
			})
		})
	}
}
