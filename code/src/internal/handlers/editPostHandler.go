package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"database/sql"
	"errors"
	"io"
	"log"
	"strconv"

	"Forum-back/pkg/services"
	"net/http"
)

func EditPostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if r.Method == http.MethodGet {
		handleGetMethodPostEdit(w, r, db, session, header)
		return
	}

	if r.Method == http.MethodPost {
		handlePostMethodPostEdit(w, r, db, session, header)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

}

func handleGetMethodPostEdit(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	ps := services.NewPostService(db)
	us := services.NewUserService(db)
	rs := services.NewReactionService(db)
	cs := services.NewCategoryService(db)

	user := us.FindById(session.User_ID)
	if user == nil {
		return
	}

	post, err := fetchPost(w, r, header, ps, cs)
	if err != nil {
		return
	}

	authorized := post.User_ID == user.ID
	if !authorized {
		ShowError403(w, header)
		return
	}

	postCategories, err := cs.FindByPostId(post)
	if err != nil {
		ShowCustomError500(w, header, "Error while retrieving categories for post: "+err.Error())
		return
	}
	post.Categories = *postCategories

	categories := cs.FindAll()
	if categories == nil {
		ShowCustomError500(w, header, "No categories found. Contact the administrator to add categories.")
		return
	}

	data := dtos.EditPostPageDto{
		Header:       *header,
		Post:         *post,
		PostedDate:   post.CreatedAt.Format("January 2, 2006"),
		Categories:   *categories,
		UserReaction: rs.FindByPostAndUserId(post.ID, user.ID),
		Like:         rs.GetLikeReactionCountOnPost(post),
		Dislike:      rs.GetDislikeReactionCountOnPost(post),
	}

	tmpl, err := templates.GetTemplateWithLayout(header, "editPost", "internal/templates/publicationEdit.gohtml")
	if err != nil {
		log.Println(err)
		ShowTemplateError500(w, &dtos.HeaderDto{})
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handlePostMethodPostEdit(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	ps := services.NewPostService(db)
	us := services.NewUserService(db)
	cs := services.NewCategoryService(db)

	user := us.FindById(session.User_ID)
	if user == nil {
		return
	}

	post, ok := getPostFromBody(w, r, ps, header.IsConnected)
	if !ok {
		return
	}
	authorized := post.User_ID == user.ID
	if !authorized {
		ShowError403(w, header)
		return
	}

	if err := updatePost(w, r, header, post, ps, cs); err != nil {
		return
	}
	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusFound)
}

func updatePost(w http.ResponseWriter, r *http.Request, header *dtos.HeaderDto, post *models.Post, ps *services.PostService, cs *services.CategoryService) error {
	if err := r.ParseForm(); err != nil {
		ShowError400(w, &dtos.HeaderDto{})
		return err
	}
	if r.FormValue("title") != "" && post.Title != r.FormValue("title") {
		if len(r.FormValue("title")) > 150 {
			ShowCustomError400(w, &dtos.HeaderDto{}, "Title is too long. Maximum length is 150 characters.")
			return nil
		}
		post.Title = r.FormValue("title")
	}
	if r.FormValue("content") != "" && post.Content != r.FormValue("content") {
		post.Content = r.FormValue("content")
	}

	if imageFile, fileMeta, err := r.FormFile("image"); err == nil {
		defer imageFile.Close()
		if fileMeta.Size > 20*1024*1024 { // 20 MB
			ShowCustomError400(w, header, "Image size exceeds the maximum limit of 20 MB.")
			return nil
		}
		if imageBytes, err := io.ReadAll(imageFile); err == nil {
			post.Picture = imageBytes
		}
	}

	err := ps.UpdateCategoryFromList(r.Form["categories"], post)
	if err.Code != http.StatusOK {
		err.Header = *header
		ShowError(w, err)
		return errors.New(err.Details)
	}

	if err := ps.UpdatePost(post); err != nil {
		ShowCustomError500(w, &dtos.HeaderDto{}, "Error while updating post: "+err.Error())
		return err
	}

	return nil
}
