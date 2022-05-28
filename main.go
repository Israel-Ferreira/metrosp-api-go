package main

import "github.com/Israel-Ferreira/metrosp-api/server"

func main() {
	router := server.SetupServer(nil, nil)
	router.Run(":5000")
}
