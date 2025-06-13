package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func PromoteUser(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	userService := services.NewUserService(db)
	roleService := services.NewRoleService(db)

	userId := r.FormValue("user_id")
	if userId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	userIdParsed, err := uuid.Parse(userId)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	user := userService.FindById(userIdParsed)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	role := roleService.GetMidPermRole()
	if role == nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}
	user.Role_ID = role.ID
	user.Role = *role
	if ok := userService.Update(user); !ok {
		http.Error(w, "Failed to promote user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func DemoteUser(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	userService := services.NewUserService(db)
	roleService := services.NewRoleService(db)

	userId := r.FormValue("user_id")
	if userId == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}
	userIdParsed, err := uuid.Parse(userId)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	user := userService.FindById(userIdParsed)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	role := roleService.GetDefaultRole()
	if role == nil {
		http.Error(w, "Role not found", http.StatusNotFound)
		return
	}
	user.Role_ID = role.ID
	user.Role = *role
	if ok := userService.Update(user); !ok {
		http.Error(w, "Failed to demote user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func AdminSearchUserHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	query := r.URL.Query().Get("search")
	if query == "" {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	userService := services.NewUserService(db)
	postService := services.NewPostService(db)
	categoryService := services.NewCategoryService(db)
	reportService := services.NewReportService(db)

	allUsers := userService.FindMultipleByAny(query)
	WaitingPosts, _ := postService.FindWaitings()
	reports, _ := reportService.FindAll()
	allCategories := categoryService.FindAll()

	data := dtos.AdminPageDto{
		Header:         *header,
		AllUsers:       *allUsers,
		WaitingPosts:   WaitingPosts,
		Reports:        reports,
		AllCategories:  allCategories,
		UserManagement: true,
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "admin", "internal/templates/admin.gohtml")
	if err != nil {
		ShowTemplateError500(w, &data.Header)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
