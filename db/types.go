package db

type status struct {
	isActive bool
	lastFetched string
}

type websiteDB map[string]status