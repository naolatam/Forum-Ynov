package handlers

import (
	"Forum-back/internal/config"
	dtos "Forum-back/pkg/dtos/templates"
	"html/template"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
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
		ShowTemplateError500(w, &dtos.HeaderDto{})
		return
	}

	tmpl.Execute(w, data)
}

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
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
		ShowTemplateError500(w, &dtos.HeaderDto{})
		return
	}

	tmpl.Execute(w, data)
}
