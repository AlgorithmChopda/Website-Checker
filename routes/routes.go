package routes

import "github.com/gorilla/mux"

func UseRoute(prefix string, subRouterFunc func(router *mux.Router), router *mux.Router) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	subRouterFunc(subRouter)
}


func Routes() *mux.Router {
	router := mux.NewRouter()
	UseRoute("/website", WebsiteRoutes, router)
	return router
}