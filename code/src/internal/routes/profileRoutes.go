package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"

	"log"
	"net/http"
)

func initProfileRoute() {
	http.HandleFunc("/profile", handlers.ProfileHandler)

	http.HandleFunc("/me", mw.GetMethodOnly(mw.WithDBAndRequireAuthRedirect("/auth/login", handlers.MyProfileHandler)))
	http.HandleFunc("/me/delete", mw.PostMethodOnly(mw.WithDBAndRequireAuthRedirect("/auth/login", handlers.DeleteMyProfileHandler)))
	http.HandleFunc("/me/edit", mw.PostMethodOnly(mw.WithDBAndRequireAuthRedirect("/auth/login", handlers.EditMyProfileHandler)))

	log.Println("[ROUTING] Profile routes initialized")
}
