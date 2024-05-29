package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
)

func ListOportunity(ctx *gin.Context) {
	opportunities := []schema.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
