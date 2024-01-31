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
	err := json.NewDecoder(req.Body).Decode(&websites)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "Incomplete request")
		return
	}
	
	for _, url := range(websites.Data) {
		fmt.Println(url)
		db.Data.AddWebsite(url)
	}
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "OK")
}

func GetWebsiteStatus (res http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	
	// if all websites status is requested
	if url == "" {
		res.WriteHeader(http.StatusOK)
		for url, status := range(db.Data) {
			fmt.Fprintf(res, "Website: %s\n\tActive: %t\n\tLast Fetched: %s\n\n", url, status.IsActive, status.LastFetched)		
		}	
		return 
	}
	
	var status db.Status
	status, err := db.Data.GetStatus(url)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "No such website found")
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Website: %s\nActive: %t\nLast Fetched: %s", url, status.IsActive, status.LastFetched)
}