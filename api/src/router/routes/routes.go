package routes

import "net/http"

// Route defines all API routes
type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}
