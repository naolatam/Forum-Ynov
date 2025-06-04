package services

import (
	"Forum-back/pkg/repositories"
	"database/sql"
)

func checkDBConnection(db *sql.DB) bool {
	if db == nil {
		return false
	}
	err := db.Ping()
	if err != nil {
		return false
	}
	return true
}

func NewSessionService(db *sql.DB) *SessionService {
	if !checkDBConnection(db) {
		return nil
	}
	return &SessionService{
		repo:     repositories.NewSessionRepository(db),
		userRepo: repositories.NewUserRepository(db),
	}
}

func NewUserService(db *sql.DB) *UserService {
	if !checkDBConnection(db) {
		return nil
	}
	return &UserService{
		repo:        repositories.NewUserRepository(db),
		sessionRepo: repositories.NewSessionRepository(db),
	}
}
