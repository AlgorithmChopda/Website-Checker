package routes

import (
	"net/http"

	controllers "github.com/AlgorithmChopda/Website-Checker.git/controller"
	"github.com/gorilla/mux"
)

func WebsiteRoutes (router *mux.Router) {
	router.HandleFunc("", controllers.ReadWebsite).Methods(http.MethodPost)
	router.HandleFunc("/status", controllers.GetWebsiteStatus).Methods(http.MethodGet)
	router.HandleFunc("", controllers.ReadWebsite).Methods(http.MethodDelete)
}