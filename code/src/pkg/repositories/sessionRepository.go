package repositories

import (
	"Forum-back/pkg/models"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type SessionRepository struct {
	db *sql.DB
}

// FindByID retrieves a session by its ID.
func (repository *SessionRepository) FindByID(id uuid.UUID) (*models.Session, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM sessions WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var session models.Session
	if rows.Next() {
		err = rows.Scan(&session.ID, &session.ExpireAt, &session.User_ID)
		if err != nil {
			return nil, err
		}
		// Check if the session has expired
		if session.ExpireAt.Before(time.Now()) {
			session.Expired = true
		}
		return &session, nil
	}
	return nil, errors.New("session not found")
}

// FindByUserID retrieves a session by its associated user ID.
func (repository *SessionRepository) FindByUserID(userID uuid.UUID) (*models.Session, error) {
	if repository.db == nil {
		return nil, errors.New("connection to database isn't established")
	}
	rows, err :=
		repository.db.Query("SELECT * FROM sessions WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var session models.Session
	if rows.Next() {
		err = rows.Scan(&session.ID, &session.ExpireAt, &session.User_ID)
		if err != nil {
			return nil, err
		}
		// Check if the session has expired
		if session.ExpireAt.Before(time.Now()) {
			session.Expired = true
		}
		return &session, nil
	}
	return nil, errors.New("session not found")
}

// Create inserts a new session into the database.
func (repository *SessionRepository) Create(session *models.Session) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("INSERT INTO sessions (id, expireAt, user_id) VALUES (?, ?, ?)", session.ID, session.ExpireAt, session.User_ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes a session from the database by its ID.
func (repository *SessionRepository) Delete(id uuid.UUID) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM sessions WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

// DeleteExpiredSessions removes all sessions that have expired before the specified time.
func (repository *SessionRepository) DeleteExpiredSessions(before time.Time) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM sessions WHERE expireAt < ?", before)
	if err != nil {
		return err
	}
	return nil
}

// GetActiveSessionCount retrieves the count of active sessions (not expired).
func (repository *SessionRepository) GetActiveSessionCount() int {
	if repository.db == nil {
		return -1
	}
	var count int
	err := repository.db.QueryRow("SELECT COUNT(*) FROM sessions WHERE expireAt > ?", time.Now()).Scan(&count)
	if err != nil {
		return -1
	}
	return count
}
