package router

import (
	"github.com/gorilla/mux"
)

// NewRouter - Used for create a routes
func NewRouter() *mux.Router {
	routes = routesStruct{
		routeStruct{
			"UpdateRequests",
			"POST",
			"/updaterequests",
			UpdateRequests,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
