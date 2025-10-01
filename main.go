package main

import (
	"log"
	"net/http"

	"golang-prototype/config"
	"golang-prototype/database"
	"golang-prototype/routes"
)

func main() {
	// Load config
	config.LoadConfig()

	// Connect DB
	database.ConnectDB()

	// Setup routes (this returns your *http.ServeMux with /users, etc.)
	r := routes.SetupRoutes()

	// Wrap with prefix /api/v1
	prefixed := http.NewServeMux()
	prefixed.Handle("/api/v1/", http.StripPrefix("/api/v1", r))

	log.Println("🚀 Server running on :8080")
	// ✅ Use prefixed, not r
	if err := http.ListenAndServe(":8080", prefixed); err != nil {
		log.Fatal(err)
	}
}
