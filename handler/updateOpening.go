package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening id"
// @Param request body UpdateOpeningRequest true "Opening data to update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		logger.Errorf("unable to bind request for opening creation: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = request.Validate()
	if err != nil {
		logger.Errorf("validation error: %s", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		msgErr := errParamIsRequired("id", "queryParameter").Error()
		logger.Errorf("error getting id query parameter: %s", msgErr)
		SendError(ctx, http.StatusBadRequest, msgErr)
		return
	}

	opening := schemas.Opening{}

	err = db.First(&opening, id).Error
	if err != nil {
		msgErr := fmt.Sprintf("opening with id %s not found: %s", id, err.Error())
		logger.Errorf(msgErr)
		SendError(ctx, http.StatusNotFound, msgErr)
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	err = db.Save(&opening).Error
	if err != nil {
		msgErr := fmt.Sprintf("error updating opening with id %s: %s", id, err.Error())
		logger.Errorf(msgErr)
		SendError(ctx, http.StatusInternalServerError, msgErr)
		return
	}

	SendSuccess(ctx, "update-opening", opening)
}
