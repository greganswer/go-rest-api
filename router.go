package main

import (
	"github.com/julienschmidt/httprouter"
)

// NewRouter takes all the routes and handles them
// It also adds custom logging
func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle
		handle = route.HandlerFunc
		handle = Logger(handle)
		router.Handle(route.Method, route.Path, handle)
	}
	return router
}
