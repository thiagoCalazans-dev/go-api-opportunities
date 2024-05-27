package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get opportunity",
	})
}
