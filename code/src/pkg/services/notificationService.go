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

// FindByUser retrieves all notifications associated with a specific user.
func (s *NotificationService) FindByUser(user *models.User) (notifs *[]*models.Notification, err error) {
	if user == nil || user.ID == uuid.Nil {
		return nil, errors.New("user not defined or invalid")
	}
	if u, err := s.ur.FindById(user.ID); u == nil || err != nil {
		return nil, errors.New("user not found")
	}

	return s.repo.FindByUserId(user.ID)
}

// Create adds a new notification for a user.
func (s *NotificationService) Create(title, description string, fromUserId, userId uuid.UUID, postID uint32) (success bool) {
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
		Post_ID:     postID,
		FromUser_ID: fromUserId,
	}

	succes, err := s.repo.Create(notif)
	if err != nil {
		log.Println("Error creating notification:", err)
	}
	return succes
}
