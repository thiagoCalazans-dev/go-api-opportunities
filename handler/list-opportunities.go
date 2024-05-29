package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
)

// @BasePath /api/v1

// @Summary List opportunities
// @Description List all opportunities
// @Tags opportunity
// @Accept json
// @Produce json
// @Success 200 {object} ShowOpportunityResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunities [get]
func ListOportunity(ctx *gin.Context) {
	opportunities := []schema.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
