package router

import (
	"github.com/gorilla/mux"

	"../controllers"
)

func LoadRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Home).Methods("GET")
	return router
}
