package handlers

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/index.html")
}
