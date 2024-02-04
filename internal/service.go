package internal

import (
	"errors"
	"fmt"
)

type service struct {
	websiteRepo WebsiteRepository
}
type WebsiteService interface {
	readWebsite(urls websiteRequestObject)
	getWebsiteStatus(url string) (status, error)
	getAllWebsiteStatus() websiteDB
	CheckWebsiteStatus()
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
			svc.websiteRepo.addWebsite(url)
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

func (svc *service) CheckWebsiteStatus() {
	svc.websiteRepo.checkAllStatus()
}
