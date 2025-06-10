package routes

import (
	"Forum-back/internal/handlers"
	"log"
	"net/http"
)

func initProfileRoute() {
	http.HandleFunc("/profile", handlers.ProfileHandler)

	http.HandleFunc("/me", handlers.MyProfileHandler)
	http.HandleFunc("/me/delete", handlers.DeleteMyProfileHandler)
	http.HandleFunc("/me/edit", handlers.EditMyProfileHandler)

	log.Println("[ROUTING] Profile routes initialized")
}
