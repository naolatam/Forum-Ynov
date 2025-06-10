package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"html/template"
	"net/http"
)

func ShowError(w http.ResponseWriter, errorData dtos.ErrorPageDto) {
	tmpl, err := templates.GetTemplateWithLayout(&errorData.Header, "error", "internal/templates/error.gohtml")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(errorData.Code)
	tmpl.Execute(w, errorData)
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/templates/error.gohtml")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := dtos.ErrorPageDto{
		Code:    http.StatusNotFound,
		Message: "Not Found",
		Details: "The requested resource could not be found.",
	}

	w.WriteHeader(http.StatusNotFound)
	tmpl.Execute(w, data)
}

func ShowError404(w http.ResponseWriter, header *dtos.HeaderDto) {

	data := dtos.ErrorPageDto{
		Code:    http.StatusNotFound,
		Message: "Not Found",
		Details: "The requested resource could not be found.",
		Header:  *header,
	}
	ShowError(w, data)
}
