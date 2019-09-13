package main

import (
	"./configs"
	"./router"
	"log"
	"net/http"
)

var config configs.Conf

func main() {
	config.GetConf()

	route := router.LoadRoutes()

	//TODO handle configs through json
	server := &http.Server{
		Addr:    config.ServerHost + ":" + config.ServerPort,
		Handler: route,
	}
	log.Fatal(server.ListenAndServe())
}
