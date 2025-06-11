package handlers

import (
	dtos "Forum-back/pkg/dtos/templates"
	"net/http"
	"os"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "internal/templates/admin.gohtml"

	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		errorData := dtos.ErrorPageDto{
			Code:    http.StatusNotFound,
			Message: "Not found",
			Details: "If you are part of the support please check the path of the : " + templatePath,
		}
		ShowError(w, errorData)
		return
	}

	http.ServeFile(w, r, templatePath)
}
