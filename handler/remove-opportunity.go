package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
)

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
