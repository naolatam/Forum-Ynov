package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"database/sql"

	"Forum-back/pkg/services"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

func SearchPostsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	categoryService := services.NewCategoryService(db)
	postService := services.NewPostService(db)

	categories := categoryService.FindAll()
	if categories == nil {
		ShowCustomError500(w, header, "No categories found. Contact the administrator to add categories.")
		return
	}

	searchTerm, searchCategory, err := parseSearchParams(r)
	if err != nil {
		ShowError400(w, header)
		return
	}

	posts, err := postService.FindPostByQueryAndCategory(searchTerm, searchCategory)
	if err != nil {
		ShowCustomError500(w, header, "Error while searching posts: "+err.Error())
		return
	}

	data := dtos.SearchPostsDto{
		Header:         *header,
		Categories:     *categories,
		SearchTerm:     searchTerm,
		SearchCategory: *searchCategory,
		Posts:          *posts,
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "searchPosts", "internal/templates/findPublication.gohtml")
	if err != nil {
		ShowTemplateError500(w, &data.Header)
		return
	}

	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func NotForNowHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This feature is not implemented yet."))

	// What to do here ?
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "internal/templates/publicationEdit.gohtml")

	db, err := config.OpenDBConnection()
	if err != nil {
		ShowDatabaseError500(w, &dtos.HeaderDto{})
		return
	}
	defer db.Close()

	tmpl, err := template.ParseFiles("internal/templates/publicationEdit.gohtml", "internal/templates/components/headerComponent.gohtml")
	if err != nil {
		ShowTemplateError500(w, &dtos.HeaderDto{})
		return
	}

	tmpl.Execute(w, nil)
}

func parseSearchParams(r *http.Request) (string, *uuid.UUID, error) {
	query := r.URL.Query()
	searchTerm := query.Get("search")
	categoryStr := query.Get("category")
	if categoryStr == "" {
		return searchTerm, &uuid.Nil, nil
	}
	categoryUUID, err := uuid.Parse(categoryStr)
	if err != nil {
		return "", nil, err
	}
	return searchTerm, &categoryUUID, nil
}
