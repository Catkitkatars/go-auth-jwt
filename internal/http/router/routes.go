package router

import (
	"authjwt/internal/http/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}

func GetRoutes() []Route {
	ah := handlers.NewAuthHandler()

	return []Route{
		{
			Method:  "POST",
			Path:    "/sayHello",
			Handler: handlers.Wrap(ah.SayHello),
		},
	}
}
