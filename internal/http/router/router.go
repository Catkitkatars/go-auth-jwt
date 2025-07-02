package router

import (
	"github.com/julienschmidt/httprouter"
)

const (
	get  = "GET"
	post = "POST"
)

func InitRouter() *httprouter.Router {
	r := httprouter.New()

	routes := GetRoutes()

	for _, route := range routes {

		switch route.Method {
		case get:
			r.GET(route.Path, route.Handler)
		case post:
			r.POST(route.Path, route.Handler)
		}
	}

	return r
}
