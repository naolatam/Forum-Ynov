package routes

import (
	"log"
	"net/http"
)

func initAuthRoutes() {
	// Routes for classic authentication
	http.HandleFunc("/auth/login", nil)
	http.HandleFunc("/auth/register", nil)
	http.HandleFunc("/auth/logout", nil)

	// Routes for social authentication
	http.HandleFunc("/auth/google", nil)
	http.HandleFunc("/auth/google/callback", nil)
	http.HandleFunc("/auth/github", nil)
	http.HandleFunc("/auth/github/callback", nil)

	log.Println("[ROUTING] Auth routes initialized")

}
