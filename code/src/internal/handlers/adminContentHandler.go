package handlers

import (
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"net/http"
)

// AdminValidateContentHandler handles the validation of content (posts or comments) by an admin.
func AdminValidateContentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	postService := services.NewPostService(db)

	post, ok := getPostFromBody(w, r, postService, header.IsConnected)
	if !ok {
		return
	}
	post.Validated = true
	if err := postService.UpdatePost(post); err != nil {
		ShowCustomError500(w, header, "Failed to validate post.")
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// AdminDeleteContentHandler handles the deletion of content (posts or comments) by an admin.
func AdminDeleteContentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	postService := services.NewPostService(db)

	post, ok := getPostFromBody(w, r, postService, header.IsConnected)
	if !ok {
		return
	}
	if err := postService.Delete(post); err != nil {
		ShowCustomError500(w, header, "Failed to delete post.")
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
