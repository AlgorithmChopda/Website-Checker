package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WebsiteRequestObject struct {
	Data []string `json:"data"`
}

func ReadWebsiteHandler(websiteSvc WebsiteService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// read request
		var req WebsiteRequestObject
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Incomplete Request")
			return
		}

		// add website to DB
		websiteSvc.ReadWebsite(req)

		// response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Website added successfully")
	}
}

func GetWebsiteStatus(websiteSvc WebsiteService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		// if all websites status is requested
		if url == "" {
			w.Header().Set("Content-Type", "application/json")

			allWebsiteStatus := websiteSvc.GetAllWebsiteStatus()

			err := json.NewEncoder(w).Encode(allWebsiteStatus)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(
					struct {
						message string
					}{"Error occured"})
			}
		} else {
			websiteStatus, err := websiteSvc.GetWebsiteStatus(url)
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
		}
	}
}
