package handlers

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"net/http"
	"os"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Handle user login
	// This function will process the login request and return a response
	// Load template
	if r.Method == http.MethodPost {
		// Process login form submission
		// Validate user credentials and set session
		w.Write([]byte("Login successful"))
	} else {
		// Render login form
		http.ServeFile(w, r, "internal/templates/authentification.html")
	}
}

func setSessionCookie(
	w http.ResponseWriter,
	expireAt time.Time,
	sessionService *services.SessionService,
	user *models.User) {
	session := sessionService.FindByUser(user)

	if session == nil {
		session = sessionService.CreateWithUser(user, expireAt)
	}

	sessionCookie := &http.Cookie{
		Name:     os.Getenv("SESSION_COOKIE_NAME"),
		Value:    session.ID.String(),
		Expires:  session.ExpireAt,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, sessionCookie)
}
