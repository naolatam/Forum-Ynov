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

func (repository *SessionRepository) FindByID(id *uuid.UUID) (*models.Session, error) {
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

func (repository *SessionRepository) FindByUserID(userID *uuid.UUID) (*models.Session, error) {
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

func (repository *SessionRepository) Delete(id *uuid.UUID) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM sessions WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *SessionRepository) DeleteExpiredSessions(before *time.Time) error {
	if repository.db == nil {
		return errors.New("connection to database isn't established")
	}
	_, err := repository.db.Exec("DELETE FROM sessions WHERE expireAt < ?", before)
	if err != nil {
		return err
	}
	return nil
}
