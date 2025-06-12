package middleware

import (
	"Forum-back/internal/config"
	"Forum-back/internal/handlers"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"net/http"
)

func PostMethodOnly(
	handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			handlers.ShowError405(w, &dtos.HeaderDto{IsConnected: false})
			return
		}
		handler(w, r)
	}
}

func GetMethodOnly(
	handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			handlers.ShowError405(w, &dtos.HeaderDto{IsConnected: false})
			return
		}
		handler(w, r)
	}
}

func WithDB(
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		db, err := config.OpenDBConnection()
		if err != nil {
			handlers.ShowDatabaseError500(w, &dtos.HeaderDto{})
			return
		}
		defer db.Close()
		handler(w, r, db)
	}
}

func WithAuth(
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	return func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)

		if !isConnected {
			handlers.ShowError403(w, &dtos.HeaderDto{IsConnected: false})
			return
		}

		handler(w, r, db, session, isConnected)
	}
}

func WithHeader(
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto),
) func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {

	return func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {

		userService := services.NewUserService(db)
		ns := services.NewNotificationService(db)

		header := &dtos.HeaderDto{
			IsConnected: isConnected,
		}

		user := userService.FindById(session.User_ID)
		if user != nil {
			header.IsAdmin = userService.IsAdmin(user)
			header.IsModerator = userService.IsModerator(user)
			notif, err := ns.FindByUser(user)
			if err != nil {
				handlers.ShowCustomError500(w, header, "Unable to retrieve notifications from the database: "+err.Error())
				return
			}
			header.Notifications = *notif

		}

		handler(w, r, db, session, header)
	}
}

func WithAuthRequired(
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	return func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)
		if !isConnected {
			handlers.ShowError403(w, &dtos.HeaderDto{IsConnected: false})
			return
		}

		handler(w, r, db, session, isConnected)
	}
}

func WithAuthForbidden(
	urlToRedirect string,
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	return func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)
		if isConnected {
			http.Redirect(w, r, urlToRedirect, http.StatusSeeOther)
			return
		}

		handler(w, r, db, session, isConnected)
	}
}

func WithRequiredAuthRedirect(
	urlToRedirect string,
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	return func(w http.ResponseWriter, r *http.Request, db *sql.DB) {

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)
		if !isConnected {
			http.Redirect(w, r, urlToRedirect, http.StatusSeeOther)
			return
		}

		handler(w, r, db, session, isConnected)
	}
}
