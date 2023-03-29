package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	var err error
	request := CreateOpeningRequest{}

	err = ctx.BindJSON(&request)
	if err != nil {
		logger.Errorf("unable to bind request for opening creation: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err = request.Validate(); err != nil {
		logger.Errorf("validation error: %s", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	err = db.Create(&opening).Error
	if err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	SendSuccess(ctx, "create-opening", opening)
}
