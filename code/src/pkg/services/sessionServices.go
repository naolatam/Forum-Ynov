package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"time"

	"github.com/google/uuid"
)

type SessionService struct {
	repo     *repositories.SessionRepository
	userRepo *repositories.UserRepository
}

func (service *SessionService) FindByID(id uuid.UUID) *models.Session {
	session, err := service.repo.FindByID(id)
	if err != nil {
		return nil
	}
	return session
}

func (service *SessionService) FindByUser(user *models.User) *models.Session {
	session, err := service.repo.FindByUserID(user.ID)
	if err != nil {
		return nil
	}
	return session
}

func (service *SessionService) CreateWithUser(user *models.User, expireIn time.Time) *models.Session {
	session := models.Session{
		ID:       uuid.New(),
		User_ID:  user.ID,
		ExpireAt: expireIn,
		Expired:  false,
	}
	err := service.repo.Create(&session)
	if err != nil {
	}
	user.Session = session

	return &session
}

func (service *SessionService) CreateFromScratch(userId *uuid.UUID, expireAt *time.Time) *models.Session {
	// Basic validation of userId and expireAt
	if userId == nil || expireAt == nil {
		return nil
	}
	if time.Now().After(*expireAt) {
		return nil
	}
	// Check if user exists to avoid creating a session for a non-existent user
	if u, err := service.repo.FindById(userId); err != nil || u == nil {
		return nil
	}
	// If the user already has a session that is not expired, return it
	if session, err := service.repo.FindByUserID(*userId); err == nil && session != nil && !session.Expired {
		return session
	}

	session := models.Session{
		ID:       uuid.New(),
		User_ID:  *userId,
		ExpireAt: *expireAt,
		Expired:  false,
	}
	err := service.repo.Create(&session)
	if err != nil {
		return nil
	}
	return &session
}

func (service *SessionService) Delete(session *models.Session) error {
	if session == nil {
		return nil
	}
	err := service.repo.Delete(session.ID)
	if err != nil {
		return err
	}
	return nil
}

func (service *SessionService) DeleteExpiredSessions(before time.Time) error {
	if before.IsZero() {
		return nil
	}
	err := service.repo.DeleteExpiredSessions(before)
	if err != nil {
		return err
	}
	return nil
}
