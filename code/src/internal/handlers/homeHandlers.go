package handlers

import (
	"Forum-back/pkg/dtos/templates"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "internal/templates/index.html"

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		errorData := templates.Error{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Details: "If you are part of the support please check the path of the : " + templatePath,
		}
		ErrorHandler(w, errorData)
		return
	}

	http.ServeFile(w, r, templatePath)
}
