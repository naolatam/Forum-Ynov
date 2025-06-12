package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"
	"log"
	"net/http"
)

func initAdminRoutes() {
	http.HandleFunc("/admin", mw.GetMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.AdminHandler))))))

	http.HandleFunc("/admin/user/search", mw.GetMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.AdminSearchUserHandler))))))
	http.HandleFunc("/admin/user/promote", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.PromoteUser))))))
	http.HandleFunc("/admin/user/demote", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.DemoteUser))))))

	log.Println("[ROUTING] Admin routes initialized")
}
