package routes

import (
	"Forum-back/internal/handlers"
	"log"
	"net/http"
)

func initProfileRoute() {
	http.HandleFunc("/profile", handlers.ProfileHandler)
	http.HandleFunc("/me", handlers.MyProfileHandler)

	log.Println("[ROUTING] Profile routes initialized")
}
