package handlers

import (
	"Forum-back/pkg/models"
	"net/http"
	"os"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "internal/templates/admin.html"

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		errorData := models.Error{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Details: "If you are part of the support please check the path of the : " + templatePath,
		}
		ErrorHandler(w, errorData)
		return
	}

	http.ServeFile(w, r, templatePath)
}
