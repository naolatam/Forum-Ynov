package routes

import (
	"log"
	"net/http"
)

func InitRoutes() {
	// Initialize the routes for the application
	// This function will set up all the necessary routes and handlers
	// For example, you can define routes for user registration, login, etc.

	initStaticRoute()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Handle the root path
		w.Write([]byte("Welcome to the Forum!"))
	})
	log.Println("[ROUTING] Registered route: GET /")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Handle the health check endpoint
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	log.Println("[ROUTING] Registered route: GET /health")

	log.Println("[ROUTING] Routes initialized successfully.")

}

func initStaticRoute() {
	fs := http.FileServer(http.Dir("internal/frontEnd/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("[ROUTING] Static files route initialized at /static/")
}
