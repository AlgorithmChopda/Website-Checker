package internal

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)


func readWebsite(urls websiteRequestObject) {
	for _, url := range(urls.Data) {
		fmt.Println(url)
		if !data.isPresent(url)  {
			data.addWebsite(url)
		}		
	}
}

func getWebsiteStatus(url string) (status, error) {
	if data.isPresent(url) {
		return data.getStatus(url), nil
	}

	return status{}, errors.New("No such website found")
}

func getAllWebsiteStatus() websiteDB {
	return data 
}

// util function to check status of website every 5 second
func CheckWebsiteStatus () {
	for {
		for url := range(data) {
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
	
	data.updateStatus(url, status)
 }