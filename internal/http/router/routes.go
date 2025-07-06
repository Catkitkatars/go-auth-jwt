package router

import (
	"authjwt/internal/http/handlers"
	"authjwt/internal/middlewares"
	"net/http"
)

type Middleware func(handle http.HandlerFunc) http.HandlerFunc

type Route struct {
	Method      string
	Prefix      string
	Middlewares []Middleware
	Handler     func(*http.Request) (any, error)
}
type RouteGroup struct {
	Method      string
	Prefix      string
	Middlewares []Middleware
	Routes      []Route
}

func GetRouteGroups() []RouteGroup {
	ah := handlers.NewAuthHandler()

	return []RouteGroup{
		{
			Method: "POST",
			Prefix: "/auth",
			Routes: []Route{
				{
					Prefix:  "/registration",
					Handler: ah.Registration,
				},
				{
					Prefix:  "/login",
					Handler: ah.Login,
				},
			},
		},
		{
			Method: "POST",
			Middlewares: []Middleware{
				middlewares.AuthMiddleware,
			},
			Routes: []Route{
				{
					Prefix:  "/sayHello",
					Handler: ah.SayHello,
				},
				{
					Prefix:  "/sayByeBye",
					Handler: ah.SayByeBye,
				},
			},
		},
	}
}
