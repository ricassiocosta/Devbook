package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return a configured router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
