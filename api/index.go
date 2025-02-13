package main

import (
	"books-api/api/config"
	"books-api/api/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	routes.RegisterRoutes()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
