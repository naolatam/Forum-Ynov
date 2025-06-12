package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	userService := services.NewUserService(db)
	sessionService := services.NewSessionService(db)
	postService := services.NewPostService(db)

	countUsers, _ := userService.GetUserCount()
	countPosts, _ := postService.GetPostCount()
	countOnlineUsers, _ := sessionService.GetActiveSessionCount()

	data := dtos.HomePageDto{
		Header:           *header,
		UserCount:        countUsers,
		PostCount:        countPosts,
		ActiveUsersCount: countOnlineUsers,
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "home", "internal/templates/index.gohtml")
	if err != nil {
		ShowTemplateError500(w, &data.Header)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
