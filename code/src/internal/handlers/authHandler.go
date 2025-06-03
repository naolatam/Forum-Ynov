package handlers

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
		http.ServeFile(w, r, "internal/templates/register.html")
	}
}

func LoginViaGoogleHandler(w http.ResponseWriter, r *http.Request) {
	var (
		googleOauthConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
		oauthStateString = "random"
	)
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	log.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
