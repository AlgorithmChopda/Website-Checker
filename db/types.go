package db

type Status struct {
	IsActive bool
	LastFetched string
}

type websiteDB map[string]Status