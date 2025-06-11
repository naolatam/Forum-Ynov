package routes

import (
	"Forum-back/internal/handlers"

	"log"
	"net/http"
)

func initAuthRoutes() {
	// Routes for classic authentication
	http.HandleFunc("/auth/login", handlers.LoginHandler)
	http.HandleFunc("/auth/register", handlers.RegisterHandler)
	http.HandleFunc("/auth/logout", handlers.LogoutHandler)

	// Routes for social authentication
	http.HandleFunc("/auth/google", handlers.LoginViaGoogleHandler)
	http.HandleFunc("/auth/github", handlers.LoginViaGithubHandler)

	/*
		http.HandleFunc("/auth/google/callback", nil)
		http.HandleFunc("/auth/github/callback", nil) */

	log.Println("[ROUTING] Auth routes initialized")

}
