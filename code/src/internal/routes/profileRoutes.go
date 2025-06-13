package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"

	"log"
	"net/http"
)

func initProfileRoute() {
	http.HandleFunc("/profile", mw.GetMethodOnly(mw.WithDB(mw.WithAuth(mw.WithHeader(handlers.ProfileHandler)))))

	http.HandleFunc("/me", mw.GetMethodOnly(mw.WithDB(mw.WithRequiredAuthRedirect("/auth/login", mw.WithHeader(handlers.MyProfileHandler)))))
	http.HandleFunc("/me/delete", mw.PostMethodOnly(mw.WithDB(mw.WithRequiredAuthRedirect("/auth/login", mw.WithHeader(handlers.DeleteMyProfileHandler)))))
	http.HandleFunc("/me/edit", mw.PostMethodOnly(mw.WithDB(mw.WithRequiredAuthRedirect("/auth/login", mw.WithHeader(handlers.EditMyProfileHandler)))))

	log.Println("[ROUTING] Profile routes initialized")
}
