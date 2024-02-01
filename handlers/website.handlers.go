package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AlgorithmChopda/Website-Checker.git/db"
)

//structure code in a layered format
//handler layer - controller

//^
//|
//|

//service or business layer

//^
//|
//|

//data layer - which will interact with database

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
		if !db.Data.IsPresent(url) {
			db.Data.AddWebsite(url)
		}
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
	if(!db.Data.IsPresent(url)) {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "No such website found")
		return
	}
	
	status = db.Data.GetStatus(url)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Website: %s\nActive: %t\nLast Fetched: %s", url, status.IsActive, status.LastFetched)
}

func DeleteWebsite(res http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	if(!db.Data.IsPresent(url)) {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "No such website found")
		return
	}

	db.Data.DeleteWebsite(url);
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Website deleted successfully")
}

// util functions to check if website active or not every 5 second
func CheckWebsiteStatus () {
	for {
		for url := range(db.Data) {
			go pingWebsite(url)
		}
		time.Sleep(time.Second * 5)		
	}
}

func pingWebsite(url string) {
	res, err := http.Get(url)
	var status bool
	if err != nil {
		status = false;
	}else {
		if res.StatusCode == 200 {
			status = true;
		}
	}
	
	db.Data.UpdateStatus(url, status)
 }