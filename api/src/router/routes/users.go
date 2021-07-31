package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.Create,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.Search,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodGet,
		Function:    controllers.Show,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodPut,
		Function:    controllers.Update,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.Delete,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}/follow",
		Method:      http.MethodPost,
		Function:    controllers.Follow,
		RequireAuth: true,
	},
}
