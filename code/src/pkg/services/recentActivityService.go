package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type RecentActivityService struct {
	repo *repositories.RecentActivityRepository
	ur   *repositories.UserRepository
	pr   *repositories.PostRepository
}

func (s *RecentActivityService) FindByUser(user *models.User) (activities *[]*models.RecentActivity, err error) {
	if user == nil || user.ID == uuid.Nil {
		return nil, errors.New("user not defined or invalid")
	}
	if u, err := s.ur.FindById(user.ID); u == nil || err != nil {
		return nil, errors.New("user not found")
	}

	return s.repo.FindByUserId(user.ID)
}

func (s *RecentActivityService) Create(action, details string, subTitle *string, userId uuid.UUID, postId uint32) (success bool) {
	if userId == uuid.Nil {
		return false
	}
	if user, err := s.ur.FindById(userId); user == nil || err != nil {
		return false
	}
	if post, err := s.pr.FindById(postId); post == nil || err != nil {
		return false
	}

	activity := &models.RecentActivity{
		Action:    action,
		Details:   details,
		SubTitle:  subTitle,
		CreatedAt: time.Now(),
		User_ID:   userId,
		Post_ID:   postId,
	}

	succes, err := s.repo.Create(activity)
	if err != nil {
		log.Println("Error finding all categories:", err)
	}
	return succes
}
