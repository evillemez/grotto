package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting... !!")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", NewRouter(appRoutes)))
	fmt.Println("Exited.")
}
