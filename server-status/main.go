package main

import (
	"log"
	"net/http"

	"./router"
)

func main() {
	route := router.LoadRoutes()

	//TODO handle config through json
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: route,
	}
	log.Fatal(server.ListenAndServe())
}
