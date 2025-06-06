package routes

import (
	"Forum-back/internal/handlers"
	"log"
	"net/http"
)

func initAdminRoutes() {
	http.HandleFunc("/admin", handlers.AdminHandler)
	log.Println("[ROUTING] Home routes initialized")
}
