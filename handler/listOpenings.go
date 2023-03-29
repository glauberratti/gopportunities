package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Show openings
// @Description Show all jobs opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	err := db.Find(&openings).Error
	if err != nil {
		msgErr := fmt.Sprintf("error listing openings: %s", err.Error())
		logger.Errorf(msgErr)
		SendError(ctx, http.StatusNotFound, msgErr)
		return
	}

	SendSuccess(ctx, "list-openings", openings)
}
