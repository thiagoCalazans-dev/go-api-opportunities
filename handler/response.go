package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(gc *gin.Context, code int, msg string) {
	gc.Header("Content-type", "application/json")
	gc.JSON(code, gin.H{
		"message": msg,
	})
}

func sendSuccess(gc *gin.Context, op string, data interface{}) {
	gc.Header("Content-type", "application/json")
	gc.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successfull", op),
		"data":    data,
	})
}
