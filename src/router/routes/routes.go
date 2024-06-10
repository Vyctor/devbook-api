package routes

import (
	"devbook-api/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri                   string
	Method                string
	Function              func(w http.ResponseWriter, r *http.Request)
	RequireAuthentication bool
}

func Setup(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoute)
	for _, route := range routes {
		if route.RequireAuthentication {
			r.HandleFunc(route.Uri, middlewares.Authenticate(route.Function)).Methods(route.Method)
			continue
		} else {
			r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
		}
	}
	return r
}
