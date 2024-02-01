package internal

import "time"

func (data *websiteDB) isPresent(websiteUrl string) bool{
	// if website not present
	if _, ok := (*data)[websiteUrl] ; !ok {
		return false;
	}

	return true;
}

func (data *websiteDB) addWebsite(websiteUrl string) {
	(*data)[websiteUrl] = status{lastFetched: "not fetched yet"}
}

func (data *websiteDB) updateStatus(websiteUrl string, isActive bool) {
	(*data)[websiteUrl] = status{isActive, time.Now().Format("2006-01-02 15:04:05")}
}

func (data *websiteDB) getStatus(websiteUrl string) (status) {
	status := (*data)[websiteUrl]
	return status
}

func (data *websiteDB) deleteWebsite(websiteUrl string) {
	delete(*data, websiteUrl)
}
