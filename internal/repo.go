package internal

import (
	"time"
)

type Status struct {
	IsActive    bool   `json:"IsActive"`
	LastFetched string `json:"LastFetched"`
}
type WebsiteDB map[string]Status

var data = WebsiteDB{}

type WebsiteRepository interface {
	isPresent(websiteUrl string) bool
	addWebsite(websiteUrl string, isActive bool)
	updateStatus(websiteUrl string, isActive bool)
	getStatus(websiteUrl string) Status
	deleteWebsite(websiteUrl string)
	getAllStatus() WebsiteDB
	getWebsites() []string
}

func NewRespository() WebsiteRepository {
	return &WebsiteDB{}
}

func (data *WebsiteDB) isPresent(websiteUrl string) bool {
	// if website not present
	if _, ok := (*data)[websiteUrl]; !ok {
		return false
	}
	return true
}

func (data *WebsiteDB) getStatus(websiteUrl string) Status {
	status := (*data)[websiteUrl]
	return status
}

func (data *WebsiteDB) getAllStatus() WebsiteDB {
	dataMap := WebsiteDB{}
	for key, val := range *data {
		dataMap[key] = val
	}
	return dataMap
}

func (data *WebsiteDB) getWebsites() []string {
	var urls []string
	for key := range *data {
		urls = append(urls, key)
	}
	return urls
}

func (data *WebsiteDB) addWebsite(websiteUrl string, isActive bool) {
	(*data)[websiteUrl] = Status{isActive, time.Now().Format("2006-01-02 15:04:05")}
}

func (data *WebsiteDB) updateStatus(websiteUrl string, isActive bool) {
	(*data)[websiteUrl] = Status{isActive, time.Now().Format("2006-01-02 15:04:05")}
}

func (data *WebsiteDB) deleteWebsite(websiteUrl string) {
	delete(*data, websiteUrl)
}
