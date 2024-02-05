package internal

import (
	"time"
)

type status struct {
	IsActive    bool   `json:"IsActive"`
	LastFetched string `json:"LastFetched"`
}
type websiteDB map[string]status

var data = websiteDB{}

type WebsiteRepository interface {
	isPresent(websiteUrl string) bool
	addWebsite(websiteUrl string, isActive bool)
	updateStatus(websiteUrl string, isActive bool)
	getStatus(websiteUrl string) status
	deleteWebsite(websiteUrl string)
	getAllStatus() websiteDB
	getWebsites() []string
}

func NewRespository() WebsiteRepository {
	return &websiteDB{}
}

func (data *websiteDB) isPresent(websiteUrl string) bool {
	// if website not present
	if _, ok := (*data)[websiteUrl]; !ok {
		return false
	}
	return true
}

func (data *websiteDB) getStatus(websiteUrl string) status {
	status := (*data)[websiteUrl]
	return status
}

func (data *websiteDB) getAllStatus() websiteDB {
	dataMap := websiteDB{}
	for key, val := range *data {
		dataMap[key] = val
	}
	return dataMap
}

func (data *websiteDB) getWebsites() []string {
	var urls []string
	for key := range *data {
		urls = append(urls, key)
	}
	return urls
}

func (data *websiteDB) addWebsite(websiteUrl string, isActive bool) {
	(*data)[websiteUrl] = status{isActive, time.Now().Format("2006-01-02 15:04:05")}
}

func (data *websiteDB) updateStatus(websiteUrl string, isActive bool) {
	(*data)[websiteUrl] = status{isActive, time.Now().Format("2006-01-02 15:04:05")}
}

func (data *websiteDB) deleteWebsite(websiteUrl string) {
	delete(*data, websiteUrl)
}
