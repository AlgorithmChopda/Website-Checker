package main

import (
	"fmt"
	"net/http"

	controllers "github.com/AlgorithmChopda/Website-Checker.git/handlers"
	"github.com/AlgorithmChopda/Website-Checker.git/routes"
)

func main() {
	router := routes.Routes()
	fmt.Println("Server Started....")
	go controllers.CheckWebsiteStatus()
	http.ListenAndServe("localhost:8000", router)
}