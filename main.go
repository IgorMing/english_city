package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Printf("Running server on port %s", port)
	router := NewRouter()
	log.Fatal(http.ListenAndServe(port, router))
}
