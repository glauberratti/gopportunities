package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening id"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		msgErr := errParamIsRequired("id", "queryParameter").Error()
		logger.Errorf("error getting id query parameter: %s", msgErr)
		SendError(ctx, http.StatusBadRequest, msgErr)
		return
	}

	opening := schemas.Opening{}
	err := db.First(&opening).Error
	if err != nil {
		msgErr := fmt.Sprintf("opening with id %s not found: %s", id, err.Error())
		logger.Errorf(msgErr)
		SendError(ctx, http.StatusNotFound, msgErr)
		return
	}

	SendSuccess(ctx, "show-opening", opening)
}
