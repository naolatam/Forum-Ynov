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

	http.HandleFunc("/admin/category", mw.GetMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.AdminCategoryHandler))))))
	http.HandleFunc("/admin/category/create", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.AdminCreateNewCategoryHandler))))))
	http.HandleFunc("/admin/category/edit", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.AdminEditCategoryHandler))))))
	http.HandleFunc("/admin/category/delete", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(mw.WithHeader((handlers.AdminDeleteCategoryHandler))))))

	log.Println("[ROUTING] Admin routes initialized")
}
