package main

import (
	"github.com/Israel-Ferreira/metrosp-api/config"
	"github.com/Israel-Ferreira/metrosp-api/repo"
	"github.com/Israel-Ferreira/metrosp-api/server"
	"github.com/Israel-Ferreira/metrosp-api/services"
)

func main() {

	config.SetupDb()

	lineRepo := repo.NewLineRepo(config.Db)
	stationRepo := repo.NewStationDbRepo(config.Db)

	lineSvc := services.NewLineService(lineRepo)
	stationSvc := services.NewStationService(stationRepo, lineRepo)

	router := server.SetupServer(stationSvc, lineSvc)
	router.Run(":5000")
}
