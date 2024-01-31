package main

import (
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/Website-Checker.git/routes"
)

func main() {
	router := routes.Routes()
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	
	})
	fmt.Println("Server Started....")
	err := http.ListenAndServe("localhost:8000", router)
	fmt.Println(err)
}