package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type NotificationService struct {
	repo *repositories.NotificationRepository
	ur   *repositories.UserRepository
}

func (s *NotificationService) FindByUser(user *models.User) (notifs *[]*models.Notification, err error) {
	if user == nil || user.ID == uuid.Nil {
		return nil, errors.New("user not defined or invalid")
	}
	if u, err := s.ur.FindById(user.ID); u == nil || err != nil {
		return nil, errors.New("user not found")
	}

	return s.repo.FindByUserId(user.ID)
}

func (s *NotificationService) Create(title, description string, userId uuid.UUID) (success bool) {
	if userId == uuid.Nil {
		return false
	}
	if user, err := s.ur.FindById(userId); user == nil || err != nil {
		return false
	}

	notif := &models.Notification{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		User_ID:     userId,
	}

	succes, err := s.repo.Create(notif)
	if err != nil {
		log.Println("Error creating notification:", err)
	}
	return succes
}
