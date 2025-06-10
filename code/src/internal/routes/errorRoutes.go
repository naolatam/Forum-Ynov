package routes

import (
	"Forum-back/internal/handlers"
	"log"
	"net/http"
)

func initErrorRoutes() {
	http.HandleFunc("/error", handlers.DefaultErrorHandler)
	log.Println("[ROUTING] Error routes initialized")
}
