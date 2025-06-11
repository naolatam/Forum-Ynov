package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"

	"log"
	"net/http"
)

func initHomeRoutes() {
	http.HandleFunc("/", mw.GetMethodOnly(mw.WithDBAndAuth(handlers.HomeHandler)))
	http.HandleFunc("/home", mw.GetMethodOnly(mw.WithDBAndAuth(handlers.HomeHandler)))

	log.Println("[ROUTING] Home routes initialized")
}
