package router

import (
	"authjwt/internal/http/handlers"
	"authjwt/internal/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var Router *httprouter.Router

func InitRouter() *httprouter.Router {
	Router = httprouter.New()

	eachRouteGroups(GetRouteGroups()...)

	return Router
}

func eachRouteGroups(groups ...RouteGroup) {
	for _, group := range groups {
		if len(group.Routes) != 0 {
			for _, route := range group.Routes {
				if !checkEmptyDataRoute(route) {
					continue
				}

				path := group.Prefix + route.Prefix
				var method string
				if route.Method == "" {
					method = group.Method
					if method == "" {
						logger.Logger.Error("eachRouteGroups: Method is empty for route", path)
						continue
					}
				}
				middlewares := append(group.Middlewares, route.Middlewares...)
				handler := eachMiddlewares(handlers.Wrap(route.Handler), middlewares...)
				Router.Handler(method, path, handler)
			}
		}
	}
}

func checkEmptyDataRoute(route Route) bool {
	if route.Prefix == "" {
		logger.Logger.Error("checkEmptyDataRoute: Prefix is empty for some route")
		return false
	}
	if route.Handler == nil {
		logger.Logger.Error("checkEmptyDataRoute: Handler is nil for route", route.Prefix)
		return false
	}
	return true
}

func eachMiddlewares(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}
