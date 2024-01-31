package db

import (
	"errors"
	"time"
)

func (data *websiteDB) AddWebsite(websiteUrl string) {
	// if website not present
	if _, ok := (*data)[websiteUrl] ; !ok {
		(*data)[websiteUrl] = Status{LastFetched: "not fetched yet"}
	}
}

func (data *websiteDB) UpdateStatus(websiteUrl string, isActive bool) {
	(*data)[websiteUrl] = Status{isActive, time.Now().Format("2006-01-02 15:04:05")}
}

func (data *websiteDB) GetStatus(websiteUrl string) (Status, error) {
	if _, ok := (*data)[websiteUrl] ; !ok {
		return Status{}, errors.New("No such website found")
	}

	status := (*data)[websiteUrl]
	return status, nil
}