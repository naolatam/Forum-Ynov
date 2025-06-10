package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/services"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	db, err := config.OpenDBConnection()
	if err != nil {

		ShowDatabaseError500(w, &dtos.HeaderDto{
			IsConnected: false})
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
		ShowTemplateError500(w, &data.Header)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
