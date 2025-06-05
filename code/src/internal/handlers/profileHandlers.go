package handlers

import (
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/profile.gohtml")
}

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/profile.gohtml")
}
