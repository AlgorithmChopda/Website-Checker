package db

func (data *websiteDB) AddWebsite(websiteUrl string) {
	// if website not present
	if _, ok := (*data)[websiteUrl] ; !ok {
		(*data)[websiteUrl] = status{}
	}
}