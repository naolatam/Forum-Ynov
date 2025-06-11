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

func WithDBAndAuth(
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		db, err := config.OpenDBConnection()
		if err != nil {
			handlers.ShowDatabaseError500(w, &dtos.HeaderDto{})
			return
		}
		defer db.Close()

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)

		handler(w, r, db, session, isConnected)
	}
}

func WithDBAndAuthRequired(
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		db, err := config.OpenDBConnection()
		if err != nil {
			handlers.ShowDatabaseError500(w, &dtos.HeaderDto{})
			return
		}
		defer db.Close()

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)
		if !isConnected {
			handlers.ShowError403(w, &dtos.HeaderDto{IsConnected: false})
			return
		}

		handler(w, r, db, session, isConnected)
	}
}

func WithDBAndAuthForbidden(
	urlToRedirect string,
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		db, err := config.OpenDBConnection()
		if err != nil {
			handlers.ShowDatabaseError500(w, &dtos.HeaderDto{})
			return
		}
		defer db.Close()

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)
		if isConnected {
			http.Redirect(w, r, urlToRedirect, http.StatusSeeOther)
			return
		}

		handler(w, r, db, session, isConnected)
	}
}

func WithDBAndRequireAuthRedirect(
	urlToRedirect string,
	handler func(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		db, err := config.OpenDBConnection()
		if err != nil {
			handlers.ShowDatabaseError500(w, &dtos.HeaderDto{})
			return
		}
		defer db.Close()

		sessionService := services.NewSessionService(db)
		isConnected, session := sessionService.IsAuthenticated(r)
		if !isConnected {
			http.Redirect(w, r, urlToRedirect, http.StatusSeeOther)
			return
		}

		handler(w, r, db, session, isConnected)
	}
}
