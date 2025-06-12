package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	userService := services.NewUserService(db)
	postService := services.NewPostService(db)
	categoryService := services.NewCategoryService(db)

	allUsers, _ := userService.GetAllUsers()
	allPost, _ := postService.FindPostByQueryAndCategory("", nil)
	allCategories := categoryService.FindAll()

	data := dtos.AdminPageDto{
		Header:        *header,
		AllUsers:      allUsers,
		AllPost:       allPost,
		AllCategories: allCategories,
	}

	if data.Header.IsAdmin == false {
		ShowError403(w, &data.Header)
		return
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "admin", "internal/templates/admin.gohtml")
	if err != nil {
		ShowTemplateError500(w, &data.Header)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
