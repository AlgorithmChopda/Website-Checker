package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/Website-Checker.git/internal"
	"github.com/gorilla/mux"
)

func main() {
	context := context.Background()
	router := mux.NewRouter()
	repository := internal.NewRespository()
	service := internal.NewService(repository)

	router.HandleFunc("/website", internal.ReadWebsiteHandler(service)).Methods(http.MethodPost)
	router.HandleFunc("/website/status", internal.GetWebsiteStatus(service)).Methods(http.MethodGet)

	fmt.Println("Server Started....")
	go service.CheckWebsiteStatus(context)
	http.ListenAndServe("localhost:8000", router)
}
