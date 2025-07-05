package router

import (
	"authjwt/internal/http/handlers"
	"authjwt/internal/middlewares"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Middleware func(handle httprouter.Handle) httprouter.Handle

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
			Path:    "/registration",
			Handler: handlers.Wrap(ah.Registration),
		},
		{
			Method:  "POST",
			Path:    "/login",
			Handler: handlers.Wrap(ah.Login),
		},
		{
			Method:  "POST",
			Path:    "/sayHello",
			Handler: handlers.Wrap(ah.SayHello),
			Middlewares: []Middleware{
				middlewares.AuthMiddleware,
			},
		},
	}
}
