package handlers

import (
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"database/sql"
	"net/http"
	"strconv"
	"time"
)

func NewCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {

	commentService := services.NewCommentService(db)
	postService := services.NewPostService(db)

	postId := r.FormValue("post_id")
	if postId == "" {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Post ID is required")
		return
	}
	postIdInt, err := strconv.Atoi(postId)
	if err != nil || postIdInt <= 0 {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Invalid Post ID")
		return
	}

	post, err := postService.FindById(uint32(postIdInt))
	if err != nil || post == nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error retrieving post: "+err.Error())
		return
	}
	content := r.FormValue("content")
	if len(content) < 1 || len(content) > 200 {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Comment content must be between 1 and 200 characters")
		return
	}

	comment := &models.Comment{
		Content:   r.FormValue("content"),
		Post_id:   uint32(postIdInt),
		User_ID:   session.User_ID,
		CreatedAt: time.Now(),
		Post:      *post,
	}

	if !commentService.CreateFromModels(comment) {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error creating comment")
		return
	}
	http.Redirect(w, r, "/posts?post_id="+postId, http.StatusSeeOther)
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	commentService := services.NewCommentService(db)
	userService := services.NewUserService(db)

	comment, success := getCommentFromBody(w, r, commentService, isConnected)
	if !success {
		return
	}

	authorized := comment.User_ID == session.User_ID
	if !authorized {
		user := userService.FindById(session.User_ID)
		authorized = userService.IsAdminOrModerator(user)
	}

	if !authorized {
		ShowError403(w, &dtos.HeaderDto{IsConnected: isConnected})
		return
	}

	if !commentService.Delete(comment) {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error deleting comment")
		return
	}

	postId := r.FormValue("post_id")
	http.Redirect(w, r, "/posts?post_id="+postId, http.StatusSeeOther)
}

func LikeCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	reactionCommentHandler(w, r, "like", db, session, isConnected)
}

func DislikeCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	reactionCommentHandler(w, r, "dislike", db, session, isConnected)
}

func getCommentFromBody(w http.ResponseWriter, r *http.Request, commentService *services.CommentService, isConnected bool) (*models.Comment, bool) {
	commentId := r.FormValue("comment_id")
	if commentId == "" {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Comment ID is required")
		return nil, false
	}
	commentIdInt, err := strconv.Atoi(commentId)
	if err != nil || commentIdInt <= 0 {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Invalid Comment ID")
		return nil, false
	}

	comment, err := commentService.FindByID(uint32(commentIdInt))
	if err != nil || comment == nil {
		ShowCustomError404(w, &dtos.HeaderDto{IsConnected: isConnected}, "This comment don't exist or cannot be retrieve: "+err.Error())
		return nil, false
	}
	return comment, true
}

func switchReactionLabel(w http.ResponseWriter, isConnected bool, reac *models.Reaction, label string, reactionService *services.ReactionService) bool {

	if reac.Label == label {
		if !reactionService.Delete(reac) {
			ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error deleting reaction")
			return false
		}
		return true // Reaction deleted successfully
	}

	switch reac.Label {
	case "dislike":
		reac.Label = "like"
		if !reactionService.Update(reac) {
			ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error updating reaction")
			return false
		}
	case "like":
		reac.Label = "dislike"
		if !reactionService.Update(reac) {
			ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error updating reaction")
			return false
		}
	default:
		break // No other label exist
	}
	return true // Reaction updated successfully
}

func reactionCommentHandler(w http.ResponseWriter, r *http.Request, label string, db *sql.DB, session *models.Session, isConnected bool) {
	commentService := services.NewCommentService(db)
	userService := services.NewUserService(db)
	reactionService := services.NewReactionService(db)

	comment, ok := getCommentFromBody(w, r, commentService, isConnected)
	if !ok {
		return
	}

	reac := reactionService.FindByCommentAndUserId(comment.ID, session.User_ID)
	if reac == nil {
		user := userService.FindById(session.User_ID)
		if user == nil {
			ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error retrieving user")
			return
		}
		if err := reactionService.Create(nil, &comment.ID, user, label); err != nil {
			ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error creating reaction: "+err.Error())
			return
		}
	} else {
		if !switchReactionLabel(w, isConnected, reac, label, reactionService) {
			return
		}
	}

	http.Redirect(w, r, "/posts?post_id="+r.FormValue("post_id"), http.StatusSeeOther)
}
