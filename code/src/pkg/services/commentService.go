package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"

	"github.com/google/uuid"
)

type CommentService struct {
	repo *repositories.CommentRepository
	us   *repositories.UserRepository
	ps   *repositories.PostRepository
}

func (s *CommentService) FindByPost(post *models.Post) (*[]*models.Comment, error) {
	if post == nil {
		return nil, nil
	}
	comments, err := s.repo.FindByPostId(post.ID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindByID(commentId uint32) (*models.Comment, error) {

	comment, err := s.repo.FindById(commentId)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (service *CommentService) GetUserCommentCount(user *models.User) int {
	if user == nil || user.ID == uuid.Nil {
		return -1
	}
	count, _ := service.repo.GetUserCommentCount(&user.ID)

	return count
}

func (service *CommentService) CreateFromModels(comment *models.Comment) bool {
	if comment == nil {
		return false
	}
	if p, err := service.ps.FindById(comment.Post_id); err != nil && p == nil {
		return false // Post does not exist
	}
	if u, err := service.us.FindById(comment.User_ID); err != nil && u == nil {
		return false // User does not exist
	}

	if err := service.repo.Create(comment); err != nil {
		return false
	}

	return true
}

func (service *CommentService) Delete(comment *models.Comment) bool {
	if comment == nil {
		return false
	}
	if err := service.repo.Delete(comment); err != nil {
		return false
	}
	return true
}
