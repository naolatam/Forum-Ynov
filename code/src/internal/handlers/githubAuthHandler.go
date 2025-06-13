package handlers

import (
	"Forum-back/internal/config"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils/oauth"
	"fmt"
	"net/http"
	"time"
)

// LoginViaGithubHandler initiates the OAuth flow for GitHub login.
func LoginViaGithubHandler(w http.ResponseWriter, r *http.Request) {
	githubOauthConfig := oauth.GetGithubOauthConfig()
	state := oauth.GenerateStateOauthCookie(w)
	url := githubOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// LoginViaGithubCallbackHandler handles the callback from GitHub after user authentication.
func LoginViaGithubCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is valid, only GET is allowed
	if r.Method != http.MethodGet {
		ShowError405(w, &dtos.HeaderDto{IsConnected: false})
		return
	}

	// Check if the state is valid, if not redirect to the login page
	if !callbackCheckState(w, r) {
		return
	}
	// If the state is valid, proceed to get the user info from Google
	code := r.FormValue("code")
	userInfo, err := oauth.GetGithubUserInfoFromCode(code)

	// If there is an error getting the user info, return an error
	if err != nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: false}, fmt.Sprintf("Failed to get user info: %v", err))
		return
	}

	// If the user info is set, search the user in database
	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{IsConnected: false})
		return
	}
	defer db.Close()

	// Initialize user and session service
	userService := services.NewUserService(db)
	sessionService := services.NewSessionService(db)

	// Search for the user by Google ID
	user := userService.FindByGithubId(userInfo.ID)

	if user == nil { // User not found, search if there is a user with the same email

		var email *string
		user, email = userService.CheckUserWithSameGithubEmails(userInfo)
		if user != nil && email == nil {
			ShowCustomError500(w, &dtos.HeaderDto{IsConnected: false}, "Failed to update user with Github ID")
			return
		}
		if email == nil && user == nil {
			user, err = userService.CreateFromGithub(userInfo)
			if err != nil {
				ShowCustomError500(w, &dtos.HeaderDto{IsConnected: false}, fmt.Sprintf("Failed to create user: %v", err))
				return
			}
		}

	}

	// Set the session cookie for the user or use the existing session if there is one
	setSessionCookie(w, time.Now().Add(6*time.Hour), sessionService, user)

	// Redirect to the home page after successful login
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// callbackCheckState checks if the state parameter in the request matches the one stored in the cookie.
func callbackCheckState(w http.ResponseWriter, r *http.Request) bool {
	state := r.FormValue("state")
	stateFromCookie, err := r.Cookie("oauthstate")
	if err != nil {
		ShowError400(w, &dtos.HeaderDto{IsConnected: false})
		return false
	}
	// Check if the state parameter in the request matches the one in the cookie
	// Else redirect to the login page
	if state != stateFromCookie.Value {
		http.Redirect(w, r, "/auth/login", http.StatusUnauthorized)
		return false
	}
	return true
}
