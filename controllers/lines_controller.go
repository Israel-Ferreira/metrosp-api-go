package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/metrosp-api/data"
	"github.com/Israel-Ferreira/metrosp-api/exceptions"
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

func (controller LinesController) Create(ctx *gin.Context) {
	var lineDTO data.LineDTO

	if err := json.NewDecoder(ctx.Request.Body).Decode(&lineDTO); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})

		return
	}

	newLine, err := controller.lineService.Create(lineDTO)

	if err != nil {
		if err == exceptions.ErrorValidation {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})

		} else if err == exceptions.ErrorLineExists {
			ctx.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusCreated, newLine)

}

func (controller *LinesController) FindById(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	line, err := controller.lineService.FindById(id)

	if err != nil {
		if err == exceptions.ErrorNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"err": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, line)
}

func NewLineController(lineService services.LineService) LinesController {
	return LinesController{
		lineService: lineService,
	}
}
