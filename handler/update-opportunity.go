package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
)

func UpdateOportunity(ctx *gin.Context) {
	request := UpdateOportunityRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opportunity := schema.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opportunity not found")
		return
	}
	// Update opportunity
	if request.Role != "" {
		opportunity.Role = request.Role
	}

	if request.Company != "" {
		opportunity.Company = request.Company
	}

	if request.Location != "" {
		opportunity.Location = request.Location
	}

	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}

	if request.Link != "" {
		opportunity.Link = request.Link
	}

	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}
	// Save opportunity
	if err := db.Save(&opportunity).Error; err != nil {
		logger.Errorf("error updating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opportunity")
		return
	}
	sendSuccess(ctx, "update-opportunity", opportunity)
}
