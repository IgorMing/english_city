package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IgorMing/englishcity/database"
	"github.com/IgorMing/englishcity/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Running server on port %s\n", port)
	router := routes.NewRouter()
	if err := database.Connect(); err != nil {
		log.Println("Error connecting to the database", err)
	}
	log.Fatal(http.ListenAndServe(":"+port, router))
}
