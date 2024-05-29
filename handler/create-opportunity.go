package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiagoCalazans-dev/go-api-opportunities.git/schema"
)

func CreateOportunity(ctx *gin.Context) {
	request := CreateOportunityRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schema.Opportunity{
		Role:     request.Role,
		Company:  request.Company,
		Salary:   request.Salary,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
	}

	if err := db.Create(&opportunity).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(ctx, "create-opportunity", opportunity)
	logger.Infof("request receive: %v", request)
}
