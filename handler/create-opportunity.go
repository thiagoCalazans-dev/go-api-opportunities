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
