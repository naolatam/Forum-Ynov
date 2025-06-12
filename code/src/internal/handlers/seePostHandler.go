package handlers

import (
	"Forum-back/internal/templates"
	dtos "Forum-back/pkg/dtos/templates"
	"Forum-back/pkg/models"
	"Forum-back/pkg/utils"
	"database/sql"
	"errors"
	"log"
	"strconv"

	"Forum-back/pkg/services"
	"net/http"
)

func SeePostHandler(w http.ResponseWriter, r *http.Request, db *sql.DB, session *models.Session, header *dtos.HeaderDto) {

	ps := services.NewPostService(db)
	us := services.NewUserService(db)
	rs := services.NewReactionService(db)
	commentService := services.NewCommentService(db)
	categoryService := services.NewCategoryService(db)

	var user models.User
	if session != nil {
		user = *us.FindById(session.User_ID)
	}
	post, err := fetchPost(w, r, header, ps, categoryService)
	if err != nil {
		return
	}

	comments, success := retrieveComments(w, commentService, post, rs, us, header)
	if !success {
		return
	}

	data := dtos.PostPageDto{
		Header:       *header,
		Post:         *post,
		Comments:     comments,
		ActualUserId: user.ID,
		UserReaction: rs.FindByPostAndUserId(post.ID, user.ID),
		Like:         rs.GetLikeReactionCountOnPost(post),
		Dislike:      rs.GetDislikeReactionCountOnPost(post),
	}

	tmpl, err := templates.GetTemplateWithLayout(&data.Header, "postPage", "internal/templates/publication.gohtml")

	if err != nil {
		ShowTemplateError500(w, &dtos.HeaderDto{})
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func fetchPost(w http.ResponseWriter, r *http.Request, header *dtos.HeaderDto, ps *services.PostService, cs *services.CategoryService) (*models.Post, error) {
	postId := r.URL.Query().Get("post_id")
	if postId == "" {
		ShowError400(w, header)
		return nil, errors.New("post_id is required")
	}
	var (
		postIdInt int
		err       error
	)
	if postIdInt, err = strconv.Atoi(postId); err != nil {
		ShowError400(w, header)
		return nil, errors.New("post_id should be a valid integer")
	}
	if postIdInt <= 0 {
		ShowError400(w, header)
		return nil, errors.New("post_id should be a valid positive integer")
	}

	post, err := ps.FindById(uint32(postIdInt))
	if err != nil {
		ShowCustomError500(w, header, "Post not found or error while retrieving post: "+err.Error())
		return nil, errors.New("post not found or error while retrieving post")
	}
	if _, err = cs.FindByPostId(post); err != nil {
		ShowCustomError500(w, header, "Error while retrieving categories for post: "+err.Error())
		return nil, errors.New("post categories not found or error while retrieving categories")
	}
	if _, err = ps.FetchUserId(post); err != nil {
		ShowCustomError500(w, header, "Error while retrieving user for post: "+err.Error())
		return nil, errors.New("post owner not found or error while retrieving categories")
	}
	return post, nil
}

func translateCommentsIntoCommentsDto(
	w http.ResponseWriter,
	commentsModels *[]*models.Comment,
	rs *services.ReactionService,
	us *services.UserService,
	header *dtos.HeaderDto,
) []*dtos.CommentDto {

	var comments []*dtos.CommentDto
	for _, comment := range *commentsModels {
		// Fetch user who made the comment
		user := us.FindById(comment.User_ID)
		if user == nil {
			ShowCustomError500(w, header, "Error while retrieving user for comment")
			return nil
		}
		dateStr := utils.TimeAgo(comment.CreatedAt)

		comments = append(comments, &dtos.CommentDto{
			ID:           comment.ID,
			Content:      comment.Content,
			User:         *user,
			Date:         dateStr,
			Like:         rs.GetLikeReactionCountOnComment(comment),
			Dislike:      rs.GetDislikeReactionCountOnComment(comment),
			UserReaction: rs.FindByCommentAndUserId(comment.ID, user.ID),
		})
	}
	return comments
}

func retrieveComments(
	w http.ResponseWriter,
	commentService *services.CommentService,
	post *models.Post,
	rs *services.ReactionService,
	us *services.UserService,
	header *dtos.HeaderDto,
) (comments []*dtos.CommentDto, success bool) {
	commentsModels, err := commentService.FindByPost(post)
	if err != nil {
		ShowCustomError500(w, header, "Error while retrieving comments for post: "+err.Error())
		return nil, false
	}

	comments = translateCommentsIntoCommentsDto(w, commentsModels, rs, us, header)
	if comments == nil {
		return nil, true
	}
	return comments, true
}
