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

	eachRoutable(GetRouteGroup(), "", nil, "")

	return Router
}

func eachRoutable(routable []Routable, parentPrefix string, parentMW []Middleware, method string) {
	for _, item := range routable {
		switch r := item.(type) {
		case *Route:
			if !checkEmptyDataRoute(*r) {
				continue
			}
			path := parentPrefix + r.Prefix
			mt := r.Method
			if r.Method == "" {
				mt = method
				if method == "" {
					log.Logger.Error("eachRouteGroups: Method is empty for route", path)
					continue
				}
			}

			middlewares := append(parentMW, r.Middlewares...)
			handler := eachMiddlewares(handlers.Wrap(r.Handler), middlewares...)
			Router.Handler(mt, path, handler)
		case *RouteGroup:
			eachRoutable(r.Items, parentPrefix+r.Prefix, append(parentMW, r.Middlewares...), r.Method)
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
