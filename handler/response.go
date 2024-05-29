package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
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

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpportunityResponse struct {
	Message string                     `json:"message"`
	Data    schema.OpportunityResponse `json:"data"`
}

type RemoveOpportunityResponse struct {
	Message string                     `json:"message"`
	Data    schema.OpportunityResponse `json:"data"`
}
type ShowOpportunityResponse struct {
	Message string                     `json:"message"`
	Data    schema.OpportunityResponse `json:"data"`
}
type ListOpportunitysResponse struct {
	Message string                       `json:"message"`
	Data    []schema.OpportunityResponse `json:"data"`
}
type UpdateOpportunityResponse struct {
	Message string                     `json:"message"`
	Data    schema.OpportunityResponse `json:"data"`
}
