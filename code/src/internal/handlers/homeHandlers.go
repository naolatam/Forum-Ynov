package handlers

import (
	"Forum-back/internal/config"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/services"
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/templates/index.gohtml")
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

	sessionService := services.NewSessionService(db)
	isConnected, _ := sessionService.IsAuthenticated(r)

	data := dtos.HomePageDto{
		IsConnected: isConnected,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
		return
	}
	return
}
