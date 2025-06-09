package handlers

import "net/http"

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/admin.gohtml")
}
