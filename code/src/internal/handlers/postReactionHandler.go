package handlers

import (
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"net/http"
	"strconv"
)

func LikePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	reactionPostHandler(w, r, "like", db, session, isConnected)
}

func DisikePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	reactionPostHandler(w, r, "dislike", db, session, isConnected)
}

func getPostFromBody(w http.ResponseWriter, r *http.Request, postService *services.PostService, isConnected bool) (*models.Post, bool) {
	postId := r.FormValue("post_id")
	if postId == "" {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Post ID is required")
		return nil, false
	}
	postIdInt, err := strconv.Atoi(postId)
	if err != nil || postIdInt <= 0 {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Invalid Post ID")
		return nil, false
	}

	post, err := postService.FindById(uint32(postIdInt))
	if err != nil || post == nil {
		ShowCustomError404(w, &dtos.HeaderDto{IsConnected: isConnected}, "This post don't exist or cannot be retrieve: "+err.Error())
		return nil, false
	}
	return post, true
}

func reactionPostHandler(w http.ResponseWriter, r *http.Request, label string, db *sql.DB, session *models.Session, isConnected bool) {
	postService := services.NewPostService(db)
	userService := services.NewUserService(db)
	reactionService := services.NewReactionService(db)
	ras := services.NewRecentActivityService(db)
	ns := services.NewNotificationService(db)

	post, ok := getPostFromBody(w, r, postService, isConnected)
	if !ok {
		return
	}
	user := userService.FindById(session.User_ID)
	if user == nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error retrieving user")
		return
	}
	reac := reactionService.FindByPostAndUserId(post.ID, session.User_ID)
	if reac == nil {
		if err := reactionService.Create(&post.ID, nil, user, label); err != nil {
			ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error creating reaction: "+err.Error())
			return
		}
	} else {
		if !switchReactionLabel(w, isConnected, reac, label, reactionService) {
			return
		}
	}

	ras.Create(label+"d Post ", post.Title[:min(100, len(post.Title))], nil, session.User_ID, post.ID)
	ns.Create(label+"d post", "have leave a "+label+" on post:",
		session.User_ID, post.User_ID, post.ID)
	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusSeeOther)
}
