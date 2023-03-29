package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glauberratti/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		msgErr := errParamIsRequired("id", "queryParameter").Error()
		logger.Errorf("error getting id query parameter: %s", msgErr)
		SendError(ctx, http.StatusBadRequest, msgErr)
		return
	}

	opening := schemas.Opening{}
	err := db.First(&opening, id).Error
	if err != nil {
		msgErr := fmt.Sprintf("opening with id %s not found: %s", id, err.Error())
		logger.Errorf(msgErr)
		SendError(ctx, http.StatusNotFound, msgErr)
		return
	}

	err = db.Delete(&opening).Error
	if err != nil {
		msgErr := fmt.Sprintf("error deleting opening with id %s: %s", id, err.Error())
		logger.Errorf(msgErr)
		SendError(ctx, http.StatusNotFound, msgErr)
		return
	}

	SendSuccess(ctx, "delete-opening", opening)
}
