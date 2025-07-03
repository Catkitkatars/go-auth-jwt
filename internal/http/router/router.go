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

		handler := eachMiddlewares(route.Handler, route.Middlewares...)

		switch route.Method {
		case get:
			r.GET(route.Path, handler)
		case post:
			r.POST(route.Path, handler)
		}
	}

	return r
}

func eachMiddlewares(h httprouter.Handle, middleware ...Middleware) httprouter.Handle {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}
