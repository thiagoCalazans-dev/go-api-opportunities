package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
)

// @BasePath /api/v1

// @Summary Remove opportunity
// @Description Remove a new job opportunity
// @Tags opportunity
// @Accept json
// @Produce json
// @Param id query string true "Opportunity identification"
// @Success 200 {object} RemoveOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opportunity [delete]
func RemoveOportunity(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opportunity := schema.Opportunity{}
	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opportunity with id %s not found", id))
		return
	}

	if err := db.Delete(&opportunity).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting opportunity with id %s", id))
		return
	}
	sendSuccess(ctx, "delete-opportunity", opportunity)
}
