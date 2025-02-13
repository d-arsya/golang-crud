package api

import (
	"books-api/config"
	"books-api/routes"
	"net/http"
)

// Handler untuk Vercel Serverless Function
func Handler(w http.ResponseWriter, r *http.Request) {
	config.ConnectDB()
	routes.RegisterRoutes()
	http.DefaultServeMux.ServeHTTP(w, r)
}
