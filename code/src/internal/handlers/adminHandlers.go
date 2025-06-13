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

func AdminHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}

	userService := services.NewUserService(db)
	postService := services.NewPostService(db)
	categoryService := services.NewCategoryService(db)
	reportService := services.NewReportService(db)

	allUsers, _ := userService.GetAllUsers()
	waitingPosts, _ := postService.FindWaitings()
	reports, _ := reportService.FindAll()
	allCategories := categoryService.FindAll()

	data := dtos.AdminPageDto{
		Header:        *header,
		AllUsers:      allUsers,
		WaitingPosts:  waitingPosts,
		Reports:       reports,
		AllCategories: allCategories,
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

func AdminCategoryHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}

	userService := services.NewUserService(db)
	postService := services.NewPostService(db)
	categoryService := services.NewCategoryService(db)
	reportService := services.NewReportService(db)

	allUsers, _ := userService.GetAllUsers()
	waitingPosts, _ := postService.FindWaitings()
	reports, _ := reportService.FindAll()
	allCategories := categoryService.FindAll()

	data := dtos.AdminPageDto{
		Header:             *header,
		AllUsers:           allUsers,
		WaitingPosts:       waitingPosts,
		Reports:            reports,
		AllCategories:      allCategories,
		CategoryManagement: true,
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

func AdminReportDelete(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}

	reportService := services.NewReportService(db)

	reportIdString := r.FormValue("report_id")
	if reportIdString == "" {
		ShowCustomError400(w, header, "Report ID is required")
		return
	}
	reportId, err := uuid.Parse(reportIdString)
	if err != nil {
		ShowCustomError400(w, header, "Invalid Report ID")
		return
	}
	if err := reportService.Delete(&models.Report{ID: reportId}); err != nil {
		ShowCustomError500(w, header, "Failed to delete report: "+err.Error())
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
