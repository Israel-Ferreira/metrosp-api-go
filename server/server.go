package server

import (
	"github.com/Israel-Ferreira/metrosp-api/controllers"
	"github.com/Israel-Ferreira/metrosp-api/services"
	"github.com/gin-gonic/gin"
)

func SetupServer(stationService services.StationService, lineService services.LineService) *gin.Engine {
	router := gin.Default()

	SetupStationsRoutes(router, stationService)

	SetupLinesRoutes(router, lineService)

	return router
}

func SetupStationsRoutes(r *gin.Engine, stationService services.StationService) {

	controller := controllers.NewStationController(stationService)

	stationRouter := r.Group("/api/stations")
	{
		stationRouter.GET("", controller.GetAll)
		stationRouter.POST("", controller.Create)

		stationRouter.GET("/:id", controller.GetById)
		stationRouter.DELETE("/:id", controller.DeleteById)
		stationRouter.PUT("/:id", controller.Update)
	}
}

func SetupLinesRoutes(r *gin.Engine, lineService services.LineService) {
	lineController := controllers.NewLineController(lineService)

	lineRouter := r.Group("/api/lines")
	{
		lineRouter.GET("", lineController.GetAll)
		lineRouter.POST("", lineController.Create)

		lineRouter.GET("/:id", func(ctx *gin.Context) {})
		lineRouter.PUT("/:id", func(ctx *gin.Context) {})
		lineRouter.DELETE("/:id", func(ctx *gin.Context) {})
	}
}
