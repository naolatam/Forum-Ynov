package handlers

import (
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
)

// AdminCreateNewCategoryHandler handles the creation of a new category by an admin.
func AdminCreateNewCategoryHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	cs := services.NewCategoryService(db)

	categoryName := r.FormValue("category_name")
	if categoryName == "" {
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}
	if len(categoryName) > 50 {
		http.Error(w, "Category name exceed 50 character", http.StatusBadRequest)
		return
	}
	category := &models.Category{
		Name: categoryName,
	}
	if ok := cs.Create(category); !ok {
		ShowCustomError500(w, header, "Failed to create category.")
		return
	}

	http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
}

// AdminDeleteCategoryHandler handles the deletion of an existing category by an admin.
func AdminDeleteCategoryHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	cs := services.NewCategoryService(db)

	categoryId := r.FormValue("category_id")
	if categoryId == "" {
		ShowCustomError400(w, header, "Category id is required")
		http.Error(w, "Category id is required", http.StatusBadRequest)
		return
	}
	categoryUuid, err := uuid.Parse(categoryId)
	if err != nil || categoryUuid == uuid.Nil {
		ShowCustomError400(w, header, "Invalid category id")
		return
	}
	category := cs.FindById(categoryUuid)
	if category == nil {
		ShowCustomError404(w, header, "Category not found")
		return
	}
	if ok := cs.Delete(category); !ok {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
}

// AdminEditCategoryHandler handles the editing of an existing category by an admin.
func AdminEditCategoryHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {
	if !header.IsAdmin {
		ShowError403(w, header)
		return
	}
	cs := services.NewCategoryService(db)

	categoryId := r.FormValue("category_id")
	categoryName := r.FormValue("category_name")
	if categoryName == "" {
		ShowCustomError400(w, header, "Category name is required")
		return
	}
	if len(categoryName) > 50 {
		ShowCustomError400(w, header, "Category name exceed 50 character")
		return
	}

	if categoryId == "" {
		ShowCustomError400(w, header, "Category id is required")
		return
	}
	categoryUuid, err := uuid.Parse(categoryId)
	if err != nil || categoryUuid == uuid.Nil {
		ShowCustomError400(w, header, "Invalid category id")
		return
	}

	category := cs.FindById(categoryUuid)
	if category == nil {
		ShowCustomError404(w, header, "Category not found")
		return
	}

	category.Name = categoryName

	if ok := cs.Update(category); !ok {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/category", http.StatusSeeOther)
}
