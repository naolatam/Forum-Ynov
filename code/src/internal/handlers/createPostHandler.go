package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils"
	"database/sql"
	"io"
	"log"
	"strconv"

	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if r.Method == http.MethodGet {
		handleGetMethodPostNew(w, r, db, session, header)
		return
	}

	if r.Method == http.MethodPost {
		handlePostMethodPostNew(w, r, db, session, header)
		return
	}
}

func handleGetMethodPostNew(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	us := services.NewUserService(db)
	cs := services.NewCategoryService(db)

	user := us.FindById(session.User_ID)
	if user == nil {
		return
	}
	categories := cs.FindAll()
	if categories == nil {
		ShowCustomError500(w, header, "No categories found. Contact the administrator to add categories.")
		return
	}

	data := dtos.EditPostPageDto{
		Header:     *header,
		Categories: *categories,
		User:       *user,
		Like:       0,
		Dislike:    0,
		IsNew:      true,
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

func handlePostMethodPostNew(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	ps := services.NewPostService(db)
	us := services.NewUserService(db)
	ras := services.NewRecentActivityService(db)
	user := us.FindById(session.User_ID)
	if user == nil {
		return
	}

	post, ok := retrieveNewPostFromBody(w, r, ps, user)
	if !ok {
		return
	}
	if err := ps.Create(post); err != nil {
		ShowCustomError500(w, &dtos.HeaderDto{}, "Error while creating post: "+err.Error())
		return
	}

	err := ps.UpdateCategoryFromList(r.Form["categories"], post)
	if err.Code != http.StatusOK {
		err.Header = *header
		ShowError(w, err)
		return
	}

	ras.Create("Created a post", post.Title, nil, user.ID, post.ID)
	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusFound)
}

func retrieveNewPostFromBody(w http.ResponseWriter, r *http.Request, ps *services.PostService, user *models.User) (*models.Post, bool) {
	post := &models.Post{
		Title:     r.FormValue("title"),
		Content:   r.FormValue("content"),
		User_ID:   user.ID,
		Validated: false,
		User:      *user,
	}
	if post.Title == "" || post.Content == "" {
		ShowCustomError400(w, &dtos.HeaderDto{}, "Title and content cannot be empty.")
		return nil, false
	}

	if len(post.Title) > 150 {
		ShowCustomError400(w, &dtos.HeaderDto{}, "Title cannot exceed 150 characters.")
		return nil, false
	}

	if pictureFile, fileMeta, err := r.FormFile("image"); err == nil {
		defer pictureFile.Close()
		if fileMeta.Size > 20*1024*1024 { // 5 MB limit
			ShowCustomError400(w, &dtos.HeaderDto{}, "File size cannot exceeds 20 MB limit.")
			return nil, false
		}
		if pictureBytes, err := io.ReadAll(pictureFile); err == nil {
			post.Picture = pictureBytes
		} else {
			return nil, false
		}
	} else {
		post.Picture = utils.GetDefaultAvatar()
	}

	return post, true
}
