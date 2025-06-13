package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"database/sql"
	"strconv"

	"Forum-back/pkg/services"
	"net/http"

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

	searchTerm, searchCategory, searchFilter, err := parseSearchParams(r)
	if err != nil {
		ShowError400(w, header)
		return
	}

	posts, err := postService.FindPostByQueryAndCategory(searchTerm, searchCategory)
	if err != nil {
		ShowCustomError500(w, header, "Error while searching posts: "+err.Error())
		return
	}
	posts = postService.FilterPosts(posts, searchFilter)

	data := dtos.SearchPostsDto{
		Header:         *header,
		Categories:     *categories,
		SearchTerm:     searchTerm,
		SearchCategory: *searchCategory,
		SearchFilter:   searchFilter,
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

func DeletePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	postService := services.NewPostService(db)
	userService := services.NewUserService(db)

	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowError403(w, header)
		return
	}

	post, ok := getPostFromBody(w, r, postService, header.IsConnected)
	if !ok {
		return
	}

	authorized := post.User_ID == user.ID || userService.IsAdmin(user)
	if !authorized {
		ShowError403(w, header)
		return
	}
	if err := postService.Delete(post); err != nil {
		ShowCustomError500(w, header, "Error while deleting post: "+err.Error())
		return
	}

}

func ReportPostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsModerator && !header.IsAdmin {
		ShowError403(w, header)
		return
	}

	postService := services.NewPostService(db)
	userService := services.NewUserService(db)
	reportService := services.NewReportService(db)

	post, ok := getPostFromBody(w, r, postService, header.IsConnected)
	if !ok {
		return
	}
	user := userService.FindById(post.User_ID)
	if user == nil {
		ShowCustomError500(w, header, "Error finding user for post.")
		return
	}

	err := reportService.Create(&models.Report{
		User_id: user.ID,
		Post_id: post.ID,
	})
	if err != nil {
		ShowCustomError500(w, header, "Error while reporting post: "+err.Error())
		return
	}
	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusSeeOther)

}

func parseSearchParams(r *http.Request) (string, *uuid.UUID, string, error) {
	query := r.URL.Query()
	searchTerm := query.Get("search")
	categoryStr := query.Get("category")
	searchFilter := query.Get("filter")

	if categoryStr == "" {
		return searchTerm, &uuid.Nil, searchFilter, nil
	}
	categoryUUID, err := uuid.Parse(categoryStr)
	if err != nil {
		return searchTerm, nil, searchFilter, err
	}
	return searchTerm, &categoryUUID, searchFilter, nil
}
