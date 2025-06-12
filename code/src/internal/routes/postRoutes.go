package routes

import (
	"Forum-back/internal/handlers"
	mw "Forum-back/internal/middleware"

	"log"
	"net/http"
)

func initPostRoutes() {
	http.HandleFunc("/posts", mw.WithDB(mw.WithAuth(mw.WithHeader(handlers.SeePostHandler))))
	http.HandleFunc("/searchPosts", mw.WithDB(mw.WithAuth(mw.WithHeader(handlers.SearchPostsHandler))))
	http.HandleFunc("/posts/new", handlers.NotForNowHandler)
	http.HandleFunc("/posts/edit", handlers.EditPostHandler)
	http.HandleFunc("/posts/delete", handlers.NotForNowHandler)
	http.HandleFunc("/posts/like", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(handlers.LikePostHandler))))
	http.HandleFunc("/posts/dislike", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(handlers.DisikePostHandler))))

	http.HandleFunc("/posts/comments/add", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(handlers.NewCommentHandler))))
	http.HandleFunc("/posts/comments/edit", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(handlers.EditCommentHandler))))
	http.HandleFunc("/posts/comments/delete", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(handlers.DeleteCommentHandler))))
	http.HandleFunc("/posts/comments/like", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(handlers.LikeCommentHandler))))
	http.HandleFunc("/posts/comments/dislike", mw.PostMethodOnly(mw.WithDB(mw.WithAuthRequired(handlers.DislikeCommentHandler))))

	log.Println("[ROUTING] Posts routes initialized")
}
