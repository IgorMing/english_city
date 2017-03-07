package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IgorMing/english_city/database"
	"github.com/IgorMing/english_city/routes"
)

func main() {
	port := ":8080"
	fmt.Printf("Running server on port %s\n", port)
	router := routes.NewRouter()
	database.Connect()
	log.Fatal(http.ListenAndServe(port, router))
}
