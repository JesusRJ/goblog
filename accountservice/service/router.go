package service

import (
	"github.com/gorilla/mux"
)

// NewRouter define as rotas com base em routes (in routes.go)
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
