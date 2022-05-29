package controllers

import (
	"net/http"

	"github.com/Israel-Ferreira/metrosp-api/services"
	"github.com/gin-gonic/gin"
)

type LinesController struct {
	lineService services.LineService
}

func (controller *LinesController) GetAll(ctx *gin.Context) {
	lines, err := controller.lineService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, lines)
}

func NewLineController(lineService services.LineService) LinesController {
	return LinesController{
		lineService: lineService,
	}
}
