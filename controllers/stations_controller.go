package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Israel-Ferreira/metrosp-api/data"
	"github.com/Israel-Ferreira/metrosp-api/services"
	"github.com/gin-gonic/gin"
)

type StationsController struct {
	stationService services.StationService
}

func (sc *StationsController) GetAll(c *gin.Context) {
	stations, err := sc.stationService.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, stations)
}

func (sc *StationsController) GetById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	station, err := sc.stationService.FindById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, station)
}

func (sc *StationsController) DeleteById(c *gin.Context) {
	id, _ := c.Params.Get("id")

	if err := sc.stationService.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (sc *StationsController) Update(c *gin.Context) {
	id, _ := c.Params.Get("id")

	var stationDTO data.StationDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&stationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	if err := sc.stationService.Update(id, stationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusNoContent, nil)

}

func (sc *StationsController) Create(c *gin.Context) {
	var stationDTO data.StationDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&stationDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	station, err := sc.stationService.Create(stationDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, station)
}

func NewStationController(svc services.StationService) StationsController {
	return StationsController{
		stationService: svc,
	}
}
