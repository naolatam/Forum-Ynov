package services

import (
	"Forum-back/pkg/models"
	"Forum-back/pkg/repositories"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type SessionService struct {
	repo     *repositories.SessionRepository
	userRepo *repositories.UserRepository
}

// FindByID retrieves a session by its ID.
func (service *SessionService) FindByID(id uuid.UUID) *models.Session {
	session, err := service.repo.FindByID(id)
	if err != nil {
		return nil
	}
	return session
}

// FindByUser retrieves a session associated with a specific user.
func (service *SessionService) FindByUser(user *models.User) *models.Session {
	if user == nil || user.ID == uuid.Nil {
		return nil
	}
	session, err := service.repo.FindByUserID(user.ID)
	if err != nil {
		return nil
	}
	return session
}

// CreateWithUser creates a new session for a given user with an expiration time.
func (service *SessionService) CreateWithUser(user *models.User, expireIn time.Time) *models.Session {
	// Basic validation of user and expireIn

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

// CreateFromScratch creates a new session for a user if they don't already have an active session.
func (service *SessionService) CreateFromScratch(userId uuid.UUID, expireAt time.Time) *models.Session {
	// Basic validation of userId and expireAt
	if time.Now().After(expireAt) {
		return nil
	}
	// Check if user exists to avoid creating a session for a non-existent user
	if u, err := service.userRepo.FindById(userId); err != nil || u == nil {
		return nil
	}
	// If the user already has a session that is not expired, return it
	if session, err := service.repo.FindByUserID(userId); err == nil && session != nil && !session.Expired {
		return session
	}

	session := models.Session{
		ID:       uuid.New(),
		User_ID:  userId,
		ExpireAt: expireAt,
		Expired:  false,
	}
	err := service.repo.Create(&session)
	if err != nil {
		return nil
	}
	return &session
}

// Delete marks a session as expired and removes it from the repository.
func (service *SessionService) Delete(session *models.Session) error {
	if session == nil {
		return nil
	}
	err := service.repo.Delete(session.ID)
	if err != nil {
		return err
	}
	session.Expired = true
	session.ExpireAt = time.Now() // Mark session as expired
	return nil
}

// DeleteExpiredSessions removes all sessions that have expired before the specified time.
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

// GetActiveSessionCount retrieves the count of active sessions.
func (service *SessionService) GetActiveSessionCount() (int, error) {
	count := service.repo.GetActiveSessionCount()
	if count == -1 {
		return count, fmt.Errorf("failed to count active sessions")
	}
	return count, nil
}

// GetSessionFromRequest retrieves the session from the request using the session cookie.
func (service *SessionService) GetSessionFromRequest(r *http.Request) (*models.Session, error) {
	sessionCookie, err := r.Cookie(os.Getenv("SESSION_COOKIE_NAME"))
	if err != nil {
		return nil, err // Other error
	}

	sessionID, err := uuid.Parse(sessionCookie.Value)
	if err != nil {
		return nil, err // Invalid UUID format
	}
	session := service.FindByID(sessionID)
	if session == nil {
		return nil, nil // Session not found
	}
	return session, nil // Valid session found
}

// IsAuthenticated checks if the user is authenticated based on the session in the request.
func (service *SessionService) IsAuthenticated(r *http.Request) (bool, *models.Session) {
	session, _ := service.GetSessionFromRequest(r)
	if session == nil {
		return false, nil // No session found
	}
	if session.Expired || time.Now().After(session.ExpireAt) {
		return false, nil // Session is expired
	}
	return true, session // Session is valid and not expired
}
