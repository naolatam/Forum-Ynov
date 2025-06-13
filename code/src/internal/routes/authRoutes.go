package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"

	"log"
	"net/http"
)

func initAuthRoutes() {
	// Routes for classic authentication
	http.HandleFunc("/auth/login", mw.WithDB(mw.WithAuthForbidden("/home", handlers.LoginHandler)))
	http.HandleFunc("/auth/register", mw.WithDB(mw.WithAuthForbidden("/home", handlers.RegisterHandler)))
	http.HandleFunc("/auth/logout", mw.GetMethodOnly(mw.WithDB(mw.WithRequiredAuthRedirect("/", handlers.LogoutHandler))))

	// Routes for social authentication
	http.HandleFunc("/auth/google", mw.GetMethodOnly(handlers.LoginViaGoogleHandler))
	http.HandleFunc("/auth/github", mw.GetMethodOnly(handlers.LoginViaGithubHandler))

	// Callback routes for social authentication
	http.HandleFunc("/auth/google/callback", mw.GetMethodOnly(handlers.LoginViaGoogleCallbackHandler))
	http.HandleFunc("/auth/github/callback", mw.GetMethodOnly(handlers.LoginViaGithubCallbackHandler))

	log.Println("[ROUTING] Auth routes initialized")

}
