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

func GetRouteGroup() []Routable {
	ah := handlers.NewAuthHandler()

	return []Routable{
		&Route{
			Method: "POST",
			Prefix: "/pressure",
			Middlewares: []Middleware{
				middlewares.AuthMiddleware,
			},
			Handler: func(r *http.Request) (any, error) {
				return map[string]string{"ur-pressure": "120/70"}, nil
			},
		},
		&RouteGroup{
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
		&RouteGroup{
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
