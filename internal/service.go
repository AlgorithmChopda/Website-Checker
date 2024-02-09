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
	ReadWebsite(urls WebsiteRequestObject)
	GetWebsiteStatus(url string) (Status, error)
	GetAllWebsiteStatus() any
	PingWebsite(url string)
	CheckWebsiteStatus(ctx context.Context)
}

func NewService(websiteRepoObject WebsiteRepository) WebsiteService {
	return &service{
		websiteRepo: websiteRepoObject,
	}
}

func (svc *service) ReadWebsite(urls WebsiteRequestObject) {
	for _, url := range urls.Data {
		fmt.Println(url)
		if !svc.websiteRepo.isPresent(url) {
			// checks status of website and adds it to data
			svc.PingWebsite(url)
		}
	}
}

func (svc *service) GetWebsiteStatus(url string) (Status, error) {
	if svc.websiteRepo.isPresent(url) {
		return svc.websiteRepo.getStatus(url), nil
	}

	return Status{}, errors.New("No such website found")
}

func (svc *service) GetAllWebsiteStatus() any {
	return svc.websiteRepo.getAllStatus()
}

func (svc *service) CheckWebsiteStatus(ctx context.Context) {
	for {
		urls := svc.websiteRepo.getAllStatus()
		for url := range urls {
			go svc.PingWebsite(url)
		}
		time.Sleep(time.Second * 5)
	}
}

func (svc *service) PingWebsite(url string) {
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
