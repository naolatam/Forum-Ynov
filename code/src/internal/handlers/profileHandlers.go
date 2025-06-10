package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"html/template"
	"log"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	data := map[string]interface{}{
		/* "id":     UserID,
		"pseudo": Pseudo,
		"email":  Email, */
	}

	tmpl, err := template.ParseFiles("internal/templates/profile.gohtml", "internal/templates/components/headerComponent.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	data := map[string]interface{}{
		/* "id":     UserID,
		"pseudo": Pseudo,
		"email":  Email, */
	}

	tmpl, err := templates.GetTemplateWithLayout(&dtos.HeaderDto{}, "/me", "internal/templates/profile.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
