package db

import (
	"time"
)

func (data *websiteDB) IsPresent(websiteUrl string) bool{
	// if website not present
	if _, ok := (*data)[websiteUrl] ; !ok {
		return false;
	}

	return true;
}

func (data *websiteDB) AddWebsite(websiteUrl string) {
	(*data)[websiteUrl] = Status{LastFetched: "not fetched yet"}
}

func (data *websiteDB) UpdateStatus(websiteUrl string, isActive bool) {
	(*data)[websiteUrl] = Status{isActive, time.Now().Format("2006-01-02 15:04:05")}
}

func (data *websiteDB) GetStatus(websiteUrl string) (Status) {
	status := (*data)[websiteUrl]
	return status
}

func (data *websiteDB) DeleteWebsite(websiteUrl string) {
	delete(*data, websiteUrl)
}