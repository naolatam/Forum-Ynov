package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/services"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

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
		Header: dtos.HeaderDto{
			IsConnected: isConnected,
		},
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "home", "internal/templates/index.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
