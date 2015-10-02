package main

import 	"github.com/evillemez/grotto/handlers"

var appRoutes = Routes {
	Route {"Index", "GET", "/", handlers.Index},
	Route {"Name", "GET", "/name/{name}", handlers.Name},
	Route {"Misc", "GET", "/misc", handlers.Misc},
}
