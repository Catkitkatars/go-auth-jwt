package router

import (
	"authjwt/internal/http/handlers"
	log "authjwt/internal/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var Router *httprouter.Router

func InitRouter() *httprouter.Router {
	Router = httprouter.New()

	for _, group := range GetRouteGroup() {
		eachRouteGroups(&group, "", nil)
	}

	return Router
}

func eachRouteGroups(group *RouteGroup, parentPrefix string, parentMW []Middleware) {
	prefix := parentPrefix + group.Prefix
	mw := append(parentMW, group.Middlewares...)

	for _, item := range group.Items {
		switch r := item.(type) {
		case *Route:
			if !checkEmptyDataRoute(*r) {
				continue
			}
			path := prefix + r.Prefix
			var method string
			if r.Method == "" {
				method = group.Method
				if method == "" {
					log.Logger.Error("eachRouteGroups: Method is empty for route", path)
					continue
				}
			}
			middlewares := append(group.Middlewares, r.Middlewares...)
			handler := eachMiddlewares(handlers.Wrap(r.Handler), middlewares...)
			Router.Handler(method, path, handler)
		case *RouteGroup:
			eachRouteGroups(r, prefix, mw)
		}
	}
}

func checkEmptyDataRoute(route Route) bool {
	if route.Prefix == "" {
		log.Logger.Error("checkEmptyDataRoute: Prefix is empty for some route")
		return false
	}
	if route.Handler == nil {
		log.Logger.Error("checkEmptyDataRoute: Handler is nil for route", route.Prefix)
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
