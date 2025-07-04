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
	initAuthRoutes()
	initProfileRoute()
	initPostRoutes()
	initHomeRoutes()
	initAdminRoutes()
	initErrorRoutes()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Handle the health check endpoint
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("[ROUTING] Routes initialized successfully.")

}

func initStaticRoute() {
	fs := http.FileServer(http.Dir("internal/frontEnd/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("[ROUTING] Static files route initialized at /static/")
}
