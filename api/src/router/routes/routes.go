package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines all API routes
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// Setup insert all routes in a given router
func Setup(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, routeLogin)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
