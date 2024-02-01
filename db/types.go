package db

import "sync"

type Status struct {
	IsActive bool
	LastFetched string
}

type websiteDB map[string]Status

//Thread safe DS
type DummyData struct {
	m map[string]interface{}
	sync.RWMutex
}

func (dd *DummyData) GetFromMap(s string) interface{} {
	dd.RLock()
	val := dd.m["s"]
	dd.RUnlock()
	
	return val
}