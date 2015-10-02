package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", NewRouter()))
	fmt.Println("Exited.")
}
