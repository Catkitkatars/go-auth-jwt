package router

import (
	"authjwt/internal/http/handlers"
	"authjwt/internal/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Middleware func(httprouter.Handle) httprouter.Handle

type Route struct {
	Method      string
	Path        string
	Handler     func(http.ResponseWriter, *http.Request, httprouter.Params)
	Middlewares []Middleware
}

func GetRoutes() []Route {
	ah := handlers.NewAuthHandler()

	return []Route{
		{
			Method:  "POST",
			Path:    "/sayHello",
			Handler: handlers.Wrap(ah.SayHello),
			Middlewares: []Middleware{
				middleware.AuthMiddleware,
			},
		},
	}
}
