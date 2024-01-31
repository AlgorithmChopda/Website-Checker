package controllers

import (
	"fmt"
	"net/http"
)

func Test(res http.ResponseWriter, r *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Test API")
}

