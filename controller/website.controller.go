package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/Website-Checker.git/db"
)

type websiteRequestObject struct {
	Data []string `json:"data"`
}

func ReadWebsite(res http.ResponseWriter, req *http.Request) {
	var websites websiteRequestObject
	json.NewDecoder(req.Body).Decode(&websites)
	for _, url := range(websites.Data) {
		fmt.Println(url)
		db.Data.AddWebsite(url)
	}
	
	fmt.Println(db.Data)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "OK")
}
