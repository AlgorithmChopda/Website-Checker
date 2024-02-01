package internal

// DB types
type status struct {
	isActive bool
	lastFetched string
}
type websiteDB map[string]status
var data = websiteDB{}

type WesbiteService interface {
	readWebsite(req websiteRequestObject)
	getWebsiteStatus(url string) (status, error)
	getAllWebsiteStatus() websiteDB
}

// Request Response types
type websiteRequestObject struct {
	Data []string `json:"data"`
}