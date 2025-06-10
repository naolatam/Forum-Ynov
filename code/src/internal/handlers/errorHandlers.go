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

func ShowError400(w http.ResponseWriter, header *dtos.HeaderDto) {

	data := dtos.ErrorPageDto{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Details: "The request could not be understood by the server due to malformation.",
		Header:  *header,
	}
	ShowError(w, data)
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

func ShowError405(w http.ResponseWriter, header *dtos.HeaderDto) {

	data := dtos.ErrorPageDto{
		Code:    http.StatusMethodNotAllowed,
		Message: "Invalid Request Method",
		Details: "The requested resource cannot be accessed with the HTTP method used.",
		Header:  *header,
	}
	ShowError(w, data)
}

func ShowDatabaseError500(w http.ResponseWriter, header *dtos.HeaderDto) {

	data := dtos.ErrorPageDto{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Details: "Unable to connect to the database. If you're the administrator, please check the database connection settings.",
		Header:  *header,
	}
	ShowError(w, data)
}

func ShowTemplateError500(w http.ResponseWriter, header *dtos.HeaderDto) {

	data := dtos.ErrorPageDto{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Details: "Unable to render the requested page.",
		Header:  *header,
	}
	ShowError(w, data)
}

func ShowCustomError500(w http.ResponseWriter, header *dtos.HeaderDto, message string) {

	data := dtos.ErrorPageDto{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Details: message,
		Header:  *header,
	}
	ShowError(w, data)
}
