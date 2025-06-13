package handlers

import (
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/services"
	"Forum-back/pkg/utils"
	"database/sql"
	"net/http"
	"strconv"
	"time"
)

// NewCommentHandler handles the creation of a new comment on a post.
func NewCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	us := services.NewUserService(db)
	commentService := services.NewCommentService(db)
	postService := services.NewPostService(db)
	ras := services.NewRecentActivityService(db)
	ns := services.NewNotificationService(db)

	post, ok := getPostFromBody(w, r, postService, isConnected)
	if !ok {
		return
	}

	if pu, err := postService.FetchUserId(post); err != nil || pu == nil {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error retrieving post user")
		return
	} else {
		post.User = *pu
	}

	content := r.FormValue("content")
	if len(content) < 1 || len(content) > 200 {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Comment content must be between 1 and 200 characters")
		return
	}

	comment := &models.Comment{
		Content:   r.FormValue("content"),
		Post_id:   post.ID,
		User_ID:   session.User_ID,
		CreatedAt: time.Now(),
		Post:      *post,
	}

	if !commentService.CreateFromModels(comment) {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error creating comment")
		return
	}

	ras.Create("New comment under", post.Title[:min(100, len(post.Title))], &comment.Content, session.User_ID, post.ID)

	if user := us.FindById(session.User_ID); user != nil {
		ns.Create("New comment",
			"have leave a comment under post:",
			session.User_ID, post.User_ID, post.ID)
		mailHtmlContent := `You have a new comment on your post: 
		<a href="https://localhost:8080/posts?post_id=` + strconv.Itoa(int(post.ID)) + `">` + post.Title + `</a><br>
		Comment content: <blockquote>` + comment.Content + `</blockquote><br>
		Comment from: <a href="https://localhost:8080/users?user_id=` + user.ID.String() + `">` + user.Pseudo + `</a>`

		utils.SendHTMLNotificationEmail(post.User.Email,
			"New comment on your post: "+post.Title, mailHtmlContent)

	}

	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusSeeOther)
}

// DeleteCommentHandler handles the deletion of a comment on a post.
func DeleteCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	commentService := services.NewCommentService(db)
	userService := services.NewUserService(db)
	ras := services.NewRecentActivityService(db)

	post, success := getPostFromBody(w, r, services.NewPostService(db), isConnected)
	if !success {
		return
	}
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

	ras.Create("Deleted comment under", post.Title[:min(100, len(post.Title))], &comment.Content, session.User_ID, post.ID)

	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusSeeOther)
}

// EditCommentHandler handles the editing of an existing comment on a post.
func EditCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	commentService := services.NewCommentService(db)
	userService := services.NewUserService(db)

	post, success := getPostFromBody(w, r, services.NewPostService(db), isConnected)
	if !success {
		return
	}
	comment, success := getCommentFromBody(w, r, commentService, isConnected)
	if !success {
		return
	}
	comment.Post_id = post.ID // Ensure the comment is linked to the correct post

	authorized := comment.User_ID == session.User_ID
	if !authorized {
		user := userService.FindById(session.User_ID)
		authorized = userService.IsAdminOrModerator(user)
	}

	if !authorized {
		ShowError403(w, &dtos.HeaderDto{IsConnected: isConnected})
		return
	}

	content := r.FormValue("content")
	if len(content) < 1 || len(content) > 200 {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Comment content must be between 1 and 200 characters")
		return
	}
	comment.Content = content

	if !commentService.Update(comment) {
		ShowCustomError500(w, &dtos.HeaderDto{IsConnected: isConnected}, "Error updating comment")
		return
	}

	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusSeeOther)
}

// LikeCommentHandler and DislikeCommentHandler handle liking and disliking comments respectively.
func LikeCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	reactionCommentHandler(w, r, "like", db, session, isConnected)
}

// DislikeCommentHandler handles the disliking of a comment on a post.
func DislikeCommentHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, isConnected bool) {
	reactionCommentHandler(w, r, "dislike", db, session, isConnected)
}

// getCommentFromBody retrieves a comment from the request body and validates it.
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

// switchReactionLabel toggles the reaction label for a comment and handles deletion if the label matches.
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

// reactionCommentHandler handles the reaction (like/dislike) to a comment on a post.
func reactionCommentHandler(w http.ResponseWriter, r *http.Request, label string, db *sql.DB, session *models.Session, isConnected bool) {
	commentService := services.NewCommentService(db)
	userService := services.NewUserService(db)
	reactionService := services.NewReactionService(db)
	ras := services.NewRecentActivityService(db)

	if label != "like" && label != "dislike" {
		ShowCustomError400(w, &dtos.HeaderDto{IsConnected: isConnected}, "Invalid reaction label")
		return
	}

	post, ok := getPostFromBody(w, r, services.NewPostService(db), isConnected)
	if !ok {
		return
	}
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
	ras.Create(label+"d comment under", post.Title[:min(100, len(post.Title))], nil, session.User_ID, post.ID)

	http.Redirect(w, r, "/posts?post_id="+strconv.Itoa(int(post.ID)), http.StatusSeeOther)
}
