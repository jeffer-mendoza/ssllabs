package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"../services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	hostname := query.Get("host")

	hostJson, err := json.Marshal(services.GetInfoDomain(hostname))
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(hostJson)
}
