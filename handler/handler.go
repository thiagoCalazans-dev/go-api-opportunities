package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get opportunity",
	})
}

func ShowOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get opportunity",
	})
}

func UpdateOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get opportunity",
	})
}

func RemoveOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get opportunity",
	})
}

func ListOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get opportunity",
	})
}
