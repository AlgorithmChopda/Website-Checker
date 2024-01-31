package routes

import (
	"net/http"

	controllers "github.com/AlgorithmChopda/Website-Checker.git/controller"
	"github.com/gorilla/mux"
)

func WebsiteRoutes (router *mux.Router) {
	router.HandleFunc("", controllers.Test).Methods(http.MethodGet)
}