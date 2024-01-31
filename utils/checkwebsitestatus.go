package main

import (
	"time"
)

func checkWebsiteStatus () {
	for ;; {
		// for url, status := range(db.Data) {
		// 	go pingWebsite(url)		
		// }
		
		// go run check
		time.Sleep(time.Second * 5)		
	}
}


func pingWebsite(url string) {

}