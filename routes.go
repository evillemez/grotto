package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/evillemez/grotto/handlers"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {"Index", "GET", "/", handlers.Index},
	Route {"Name", "GET", "/name/{name}", handlers.Name},
	Route {"Misc", "GET", "/misc", handlers.Misc},
}

func NewRouter() *mux.Router {
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
