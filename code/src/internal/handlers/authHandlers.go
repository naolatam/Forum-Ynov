package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"net/http"
	"os"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Open a database connection
	// Handle any errors that may occur during the connection
	db, err := config.OpenDBConnection()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Initialize session service and check if the user is already authenticated
	sessionService := services.NewSessionService(db)
	if isConnected, _ := sessionService.IsAuthenticated(r); isConnected {
		// If the user is already authenticated, redirect to the home page
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

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
		Path:     "/",

		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, sessionCookie)
}
