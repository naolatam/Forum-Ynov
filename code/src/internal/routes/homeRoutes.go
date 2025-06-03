package routes

import (
	"Forum-back/internal/handlers"
	"log"
	"net/http"
)

func initHomeRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/home", handlers.HomeHandler)

	log.Println("[ROUTING] Home routes initialized")
}
