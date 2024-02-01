package main

import (
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/Website-Checker.git/internal"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// router.HandleFunc("/website", internal.ReadWebsiteHandler(internal.WesbiteService{})).Methods(http.MethodPost)
	// router.HandleFunc("/website/status", internal.ReadWebsiteHandler(internal.WesbiteService{})).Methods(http.MethodGet)
	
	fmt.Println("Server Started....")
	go internal.CheckWebsiteStatus()
	http.ListenAndServe("localhost:8000", router)
}