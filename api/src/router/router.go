package router

import "github.com/gorilla/mux"

// Generate will return a configured router
func Generate() *mux.Router {
	return mux.NewRouter()
}
