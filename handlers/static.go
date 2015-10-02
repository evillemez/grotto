package handlers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func Index(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(res, "Index\n")
}

func Name(res http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)
	fmt.Fprintf(res, "Greetings %s\n", vars["name"])
}

func Misc(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(res, "Misc\n")
}
