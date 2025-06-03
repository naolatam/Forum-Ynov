package routes

import (
	"Forum-back/internal/handlers"
	"log"
	"net/http"
)

func initPostRoutes() {
	http.HandleFunc("/posts", handlers.SeePostHandler)
	http.HandleFunc("/searchPosts", handlers.SearchPostsHandler)
	http.HandleFunc("/posts/new", handlers.NotForNowHandler)
	http.HandleFunc("/posts/edit", handlers.EditPostHandler)
	http.HandleFunc("/posts/delete", handlers.NotForNowHandler)
	http.HandleFunc("/posts/like", handlers.NotForNowHandler)
	http.HandleFunc("/posts/dislike", handlers.NotForNowHandler)

	log.Println("[ROUTING] Posts routes initialized")
}
