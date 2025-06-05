package handlers

import "net/http"

func SearchPostsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/findPublication.gohtml")
}

func SeePostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/publication.gohtml")
}

func NotForNowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This feature is not implemented yet."))
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/publicationEdit.gohtml")
}
