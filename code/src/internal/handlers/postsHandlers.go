package handlers

import (
	"Forum-back/internal/config"
	"net/http"
	"log"
)

func SearchPostsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/findPublication.gohtml")

	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func SeePostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/publication.gohtml")

	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func NotForNowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This feature is not implemented yet."))

	// What to do here ?
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/publicationEdit.gohtml")

	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
}
