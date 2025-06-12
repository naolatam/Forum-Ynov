package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"
	"log"
	"net/http"
)

func initAdminRoutes() {
	http.HandleFunc("/admin", mw.GetMethodOnly(mw.WithDB(mw.WithAuth(mw.WithHeader((handlers.AdminHandler))))))
	log.Println("[ROUTING] Admin routes initialized")
}
