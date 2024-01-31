package main

import (
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/Website-Checker.git/routes"
)

func main() {
	router := routes.Routes()
	fmt.Println("Server Started....")
	http.ListenAndServe("localhost:8000", router)
}