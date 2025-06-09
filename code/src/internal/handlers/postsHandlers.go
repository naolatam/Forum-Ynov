package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"

	"Forum-back/pkg/services"
	"log"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

func SearchPostsHandler(w http.ResponseWriter, r *http.Request) {
	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	categoryService := services.NewCategoryService(db)
	postService := services.NewPostService(db)
	sessionService := services.NewSessionService(db)

	isConnected, _ := sessionService.IsAuthenticated(r)

	categories := categoryService.FindAll()
	if categories == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	searchTerm, searchCategory, err := parseSearchParams(r)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	posts, err := postService.FindPostByQueryAndCategory(searchTerm, searchCategory)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := dtos.SearchPostsDto{
		Header: dtos.HeaderDto{
			IsConnected: isConnected,
		},
		Categories:     *categories,
		SearchTerm:     searchTerm,
		SearchCategory: *searchCategory,
		Posts:          *posts,
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "searchPosts", "internal/templates/findPublication.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func SeePostHandler(w http.ResponseWriter, r *http.Request) {

	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	tmpl, err := template.ParseFiles("internal/templates/publication.gohtml", "internal/templates/components/headerComponent.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)

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

	tmpl, err := template.ParseFiles("internal/templates/publicationEdit.gohtml", "internal/templates/components/headerComponent.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func parseSearchParams(r *http.Request) (string, *uuid.UUID, error) {
	query := r.URL.Query()
	searchTerm := query.Get("search")
	categoryStr := query.Get("category")
	if categoryStr == "" {
		return searchTerm, nil, nil
	}
	categoryUUID, err := uuid.Parse(categoryStr)
	if err != nil {
		return "", nil, err
	}
	return searchTerm, &categoryUUID, nil
}
