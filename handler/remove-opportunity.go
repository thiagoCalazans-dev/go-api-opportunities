package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveOportunity(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "get opportunity",
	})
}
