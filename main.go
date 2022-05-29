package main

import (
	"github.com/Israel-Ferreira/metrosp-api/config"
	"github.com/Israel-Ferreira/metrosp-api/server"
)

func main() {

	config.SetupDb()

	router := server.SetupServer(nil, nil)
	router.Run(":5000")
}
