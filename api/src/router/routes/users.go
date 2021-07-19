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
		RequireAuth: true,
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
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodPut,
		Function:    controllers.Update,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.Delete,
		RequireAuth: false,
	},
}
