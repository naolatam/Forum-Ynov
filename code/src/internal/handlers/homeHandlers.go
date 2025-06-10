package handlers

import (
	"Forum-back/internal/config"
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/services"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	db, err := config.OpenDBConnection()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	userService := services.NewUserService(db)
	sessionService := services.NewSessionService(db)
	postService := services.NewPostService(db)
	isConnected, _ := sessionService.IsAuthenticated(r)

	countUsers, _ := userService.GetUserCount() 
	countPosts, _ := postService.GetPostCount()
	countOnlineUsers, _ := sessionService.GetActiveSessionCount()
	

	data := dtos.HomePageDto{
		Header: dtos.HeaderDto{
			IsConnected: isConnected,
		},
		UserCount: countUsers,
		PostCount: countPosts,
		ActiveUsersCount: countOnlineUsers,
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "home", "internal/templates/index.gohtml")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
