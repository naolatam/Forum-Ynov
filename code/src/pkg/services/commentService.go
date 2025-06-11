package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"

	"github.com/google/uuid"
)

type CommentService struct {
	repo *repositories.CommentRepository
}

func (service *CommentService) GetUserCommentCount(user *models.User) int {
	if user == nil || user.ID == uuid.Nil {
		return -1
	}
	count, _ := service.repo.GetUserCommentCount(&user.ID)

	return count
}
