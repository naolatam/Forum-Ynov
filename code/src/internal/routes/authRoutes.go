package routes

import (
	"Forum-back/internal/handlers"

	"log"
	"net/http"
)

func initAuthRoutes() {
	// Routes for classic authentication
	http.HandleFunc("/auth/login", handlers.LoginHandler)
	/* http.HandleFunc("/auth/register", nil)
	http.HandleFunc("/auth/logout", nil)
	*/
	// Routes for social authentication
	http.HandleFunc("/auth/google", handlers.LoginViaGoogleHandler)
	http.HandleFunc("/auth/github", handlers.LoginViaGithubHandler)

	// Callback routes for social authentication
	http.HandleFunc("/auth/google/callback", handlers.LoginViaGoogleCallbackHandler)
	http.HandleFunc("/auth/github/callback", handlers.LoginViaGithubCallbackHandler)

	log.Println("[ROUTING] Auth routes initialized")

}
