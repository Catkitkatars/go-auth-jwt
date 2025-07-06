package router

import (
	"authjwt/internal/http/handlers"
	"authjwt/internal/middlewares"
	"net/http"
)

type Middleware func(handle http.HandlerFunc) http.HandlerFunc

type Routable interface {
	IsRoutable()
}
type Route struct {
	Method      string
	Prefix      string
	Middlewares []Middleware
	Handler     func(*http.Request) (any, error)
}

func (r Route) IsRoutable() {}

type RouteGroup struct {
	Method      string
	Prefix      string
	Middlewares []Middleware
	Items       []Routable
}

func (g RouteGroup) IsRoutable() {}

func GetRouteGroup() []RouteGroup {
	ah := handlers.NewAuthHandler()

	return []RouteGroup{
		{
			Method: "POST",
			Prefix: "/auth",
			Items: []Routable{
				&Route{
					Prefix:  "/registration",
					Handler: ah.Registration,
				},
				&Route{
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
			Items: []Routable{
				&Route{
					Prefix:  "/sayHello",
					Handler: ah.SayHello,
				},
				&Route{
					Prefix:  "/sayByeBye",
					Handler: ah.SayByeBye,
				},
				&RouteGroup{
					Method: "POST",
					Prefix: "/some",
					Items: []Routable{
						&Route{
							Prefix:  "/saySomeThing",
							Handler: ah.SaySomeThing,
						},
					},
				},
			},
		},
	}
}
