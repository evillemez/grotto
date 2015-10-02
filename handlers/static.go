package handlers

import (
	"fmt"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(res, "Index")
}

func Name(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(res, "Name")
}

func Misc(res http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(res, "Misc")
}
