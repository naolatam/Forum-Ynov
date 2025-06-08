package handlers

import (
	"Forum-back/internal/config"
	"html/template"
	"log"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/profile.gohtml")

	tmpl, err := template.ParseFiles("internal/templates/profile.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

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

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/profile.gohtml")

	tmpl, err := template.ParseFiles("internal/templates/profile.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	data := map[string]interface{}{
		/* "id"	: UserID,
		"pseudo": Pseudo,
		"email": Email,  */
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
