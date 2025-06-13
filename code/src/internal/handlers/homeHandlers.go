package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	userService := services.NewUserService(db)
	sessionService := services.NewSessionService(db)
	postService := services.NewPostService(db)

	countUsers, _ := userService.GetUserCount()
	countPosts, _ := postService.GetPostCount()
	countOnlineUsers, _ := sessionService.GetActiveSessionCount()

	limit := 3
	lastPosts, err := postService.FindLastPosts(&limit)
	if err != nil {
		ShowCustomError500(w, header, "Error while retrieving last posts: "+err.Error())
		return
	}

	data := dtos.HomePageDto{
		Header:           *header,
		UserCount:        countUsers,
		LastPosts:        *lastPosts,
		PostCount:        countPosts,
		ActiveUsersCount: countOnlineUsers,
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "home", "internal/templates/index.gohtml")
	if err != nil {
		log.Println("Error getting template:", err)
		ShowTemplateError500(w, &data.Header)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
