package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils/oauth"
	"fmt"
	"net/http"
	"time"
)

func LoginViaGoogleHandler(w http.ResponseWriter, r *http.Request) {
	googleOauthConfig := oauth.GetGoogleOauthConfig()
	state := oauth.GenerateStateOauthCookie(w)
	url := googleOauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func LoginViaGoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is valid, only GET is allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Check if the state is valid, if not redirect to the login page
	if !callbackCheckState(w, r) {
		return
	}
	// If the state is valid, proceed to get the user info from Google
	code := r.FormValue("code")
	userInfo, err := oauth.GetGoogleUserInfoFromCode(code)

	// If there is an error getting the user info, return an error
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user info: %v", err), http.StatusInternalServerError)
		return
	}

	// If the user info is set, search the user in database
	db, err := config.OpenDBConnection()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Initialize user and session service
	userService := services.NewUserService(db)
	sessionService := services.NewSessionService(db)

	// Search for the user by Google ID
	user := userService.FindByGoogleId(userInfo.ID)

	if user == nil {

		// User not found, search if there is a user with the same email
		user = userService.FindByEmail(userInfo.Email)
		if user != nil {
			// User with the same email found, update Google ID
			user.Google_ID = &userInfo.ID
			if !userService.Update(user) {
				http.Error(w, "Failed to update user with Google ID", http.StatusInternalServerError)
				return
			}
		} else {
			user, err = userService.CreateFromGoogle(userInfo)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
				return
			}
		}

	}
	// Set the session cookie for the user or use the existing session if there is one
	setSessionCookie(w, time.Now().Add(6*time.Hour), sessionService, user)

	// Redirect to the home page after successful login
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
