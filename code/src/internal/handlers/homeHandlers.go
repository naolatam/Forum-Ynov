package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/pkg/services"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/google/uuid"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/index.html")

	tmpl, err := template.ParseFiles("internal/templates/index.html")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	cookie, err := r.Cookie(os.Getenv("forum_session"))
	if err != nil {
		log.Println("User does not have a session cookie:", err)
	} else if cookie.Value == "" {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}
		
	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	sessionService := services.NewSessionService(db)
	sessionID := uuid.MustParse(cookie.Value)
	session := sessionService.FindByID(sessionID)
	
	data := map[string]bool{
	"isConnected": false,
	} 

	if session != nil && !session.Expired {
		data["isConnected"] = true
		}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	return
}	


