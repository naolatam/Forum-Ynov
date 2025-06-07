package handlers

import (
	"Forum-back/pkg/dtos/templates"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, errorData templates.Error) {
	tmpl, err := template.ParseFiles("internal/templates/error.gohtml")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, errorData)
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("internal/templates/error.gohtml")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := templates.Error{
		Code:    http.StatusNotFound,
		Message: "Not Found",
		Details: "The requested resource could not be found.",
	}

	w.WriteHeader(http.StatusNotFound)
	tmpl.Execute(w, data)
}
