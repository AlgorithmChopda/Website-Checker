package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type service struct {
	websiteRepo WebsiteRepository
}
type WebsiteService interface {
	readWebsite(urls websiteRequestObject)
	getWebsiteStatus(url string) (status, error)
	getAllWebsiteStatus() websiteDB
	pingWebsite(url string)
	CheckWebsiteStatus(ctx context.Context)
}

func NewService(websiteRepoObject WebsiteRepository) WebsiteService {
	return &service{
		websiteRepo: websiteRepoObject,
	}
}

func (svc *service) readWebsite(urls websiteRequestObject) {
	for _, url := range urls.Data {
		fmt.Println(url)
		if !svc.websiteRepo.isPresent(url) {
			// checks status of website and adds it to data
			svc.pingWebsite(url)
		}
	}
}

func (svc *service) getWebsiteStatus(url string) (status, error) {
	if svc.websiteRepo.isPresent(url) {
		return svc.websiteRepo.getStatus(url), nil
	}

	return status{}, errors.New("No such website found")
}

func (svc *service) getAllWebsiteStatus() websiteDB {
	return svc.websiteRepo.getAllStatus()
}

func (svc *service) CheckWebsiteStatus(ctx context.Context) {
	for {
		urls := svc.websiteRepo.getAllStatus()
		for url := range urls {
			go svc.pingWebsite(url)
		}
		time.Sleep(time.Second * 5)
	}
}

func (svc *service) pingWebsite(url string) {
	res, err := http.Get(url)
	var status bool
	if err != nil {
		status = false
	} else {
		if res.StatusCode == 200 {
			status = true
		}
	}

	svc.websiteRepo.updateStatus(url, status)
}
