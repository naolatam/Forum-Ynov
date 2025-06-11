package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"

	"log"
	"net/http"
)

func initPostRoutes() {
	http.HandleFunc("/posts", mw.WithDBAndAuth(handlers.SeePostHandler))
	http.HandleFunc("/searchPosts", mw.WithDBAndAuth(handlers.SearchPostsHandler))
	http.HandleFunc("/posts/new", handlers.NotForNowHandler)
	http.HandleFunc("/posts/edit", handlers.EditPostHandler)
	http.HandleFunc("/posts/delete", handlers.NotForNowHandler)
	http.HandleFunc("/posts/like", mw.PostMethodOnly(mw.WithDBAndAuthRequired(handlers.LikePostHandler)))
	http.HandleFunc("/posts/dislike", mw.PostMethodOnly(mw.WithDBAndAuthRequired(handlers.DisikePostHandler)))

	http.HandleFunc("/posts/comments/add", mw.PostMethodOnly(mw.WithDBAndAuthRequired(handlers.NewCommentHandler)))
	http.HandleFunc("/posts/comments/delete", mw.PostMethodOnly(mw.WithDBAndAuthRequired(handlers.DeleteCommentHandler)))
	http.HandleFunc("/posts/comments/like", mw.PostMethodOnly(mw.WithDBAndAuthRequired(handlers.LikeCommentHandler)))
	http.HandleFunc("/posts/comments/dislike", mw.PostMethodOnly(mw.WithDBAndAuthRequired(handlers.DislikeCommentHandler)))

	log.Println("[ROUTING] Posts routes initialized")
}
