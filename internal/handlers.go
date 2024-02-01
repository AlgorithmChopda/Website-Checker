package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReadWebsiteHandler(websiteSvc WesbiteService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// read request
		var req websiteRequestObject
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Incomplete Request")
			return
		}

		// add website to DB
		websiteSvc.readWebsite(req) 

		// response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Website added successfully")
	}
}

func GetWebsiteStatus(websiteSvc WesbiteService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {		
		url := r.URL.Query().Get("url")
	
		// if all websites status is requested
		if url == "" {	
			allWebsiteStatus := websiteSvc.getAllWebsiteStatus()
			
			err := json.NewEncoder(w).Encode(allWebsiteStatus)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "cannot encode data into json")
				return
			}

			// success
			w.WriteHeader(http.StatusOK)
		} else {
			websiteStatus, err := websiteSvc.getWebsiteStatus(url)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "No such URL found")	
				return 
			}
			
			err = json.NewEncoder(w).Encode(websiteStatus)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "cannot encode data into json")
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	}
}