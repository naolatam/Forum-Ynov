package handlers

import (
	"Forum-back/internal/config"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils"
	"database/sql"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
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
		postLoginHandler(w, r, db)
	} else {
		// Render login form
		showLoginPage(w, r, dtos.ErrorPageDto{})
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenDBConnection()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sessionService := services.NewSessionService(db)
	userService := services.NewUserService(db)

	if isConnected, _ := sessionService.IsAuthenticated(r); isConnected {
		// If the user is already authenticated, redirect to the home page
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm")

	if username == "" || password == "" || email == "" || confirmPassword == "" {
		showLoginPage(w, r, dtos.ErrorPageDto{Details: "Incomplete form", Message: "Username, email, and password are required"})
		return
	}

	if user := userService.FindByUsername(username); user != nil {
		showLoginPage(w, r, dtos.ErrorPageDto{Details: "User already exist", Message: "This username is already used"})
		return
	}
	if user := userService.FindByEmail(email); user != nil {
		showLoginPage(w, r, dtos.ErrorPageDto{Details: "User already exist", Message: "This email is already used"})
		return
	}

	newUser := &models.User{
		Pseudo:    username,
		Email:     email,
		CreatedAt: time.Now(),
		Avatar:    utils.GetDefaultAvatar(),
		Bio:       "",
	}
	if newUser.Password, err = utils.CheckForNewPassword(password, confirmPassword); err != nil {
		showLoginPage(w, r, dtos.ErrorPageDto{Details: "Password error", Message: err.Error()})
		return
	}

	if newUser, err = userService.Create(newUser); err != nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: false}, "Unable to create user: "+err.Error())
		return
	}

	setSessionCookie(w, time.Now().Add(6*time.Hour), services.NewSessionService(db), newUser)
	http.Redirect(w, r, "/", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Open a database connection
	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
		return
	}
	defer db.Close()

	// Initialize session service and check if the user is authenticated
	sessionService := services.NewSessionService(db)
	isConnected, session := sessionService.IsAuthenticated(r)
	if !isConnected {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Delete the session from db and cookie
	deleteSessionCookie(w)
	sessionService.Delete(session)
	// Redirect to the login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func postLoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userService := services.NewUserService(db)
	usernameOrEmail := r.FormValue("username")
	password := r.FormValue("password")
	if usernameOrEmail == "" || password == "" {
		ShowCustomError400(w, &dtos.HeaderDto{}, "Username and password are required")
		return
	}

	user := userService.FindByUsername(usernameOrEmail)
	if user == nil {
		user = userService.FindByEmail(usernameOrEmail)
		if user == nil {
			showLoginPage(w, r, dtos.ErrorPageDto{
				Message: "No user with this credentials has been found in database",
				Details: "Invalid credentials",
				Code:    http.StatusForbidden,
			})
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		showLoginPage(w, r, dtos.ErrorPageDto{
			Message: "No user with this credentials has been found in database",
			Details: "Invalid credentials",
			Code:    http.StatusForbidden,
		})
	}
	setSessionCookie(w, time.Now().Add(6*time.Hour), services.NewSessionService(db), user)
	http.Redirect(w, r, "/", http.StatusFound)
}

func showLoginPage(w http.ResponseWriter, r *http.Request, errDto dtos.ErrorPageDto) {
	tmpl, err := template.ParseFiles("internal/templates/authentification.gohtml")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := dtos.AuthenticationPageDto{
		IsRegister: r.URL.Query().Get("isRegister") == "true" || strings.Contains(r.URL.Path, "register"),
		Error:      errDto,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func setSessionCookie(
	w http.ResponseWriter,
	expireAt time.Time,
	sessionService *services.SessionService,
	user *models.User) {
	session := sessionService.FindByUser(user)

	if session == nil || session.Expired {
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

func deleteSessionCookie(w http.ResponseWriter) {

	sessionCookie := &http.Cookie{
		Name:     os.Getenv("SESSION_COOKIE_NAME"),
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, sessionCookie)
}
