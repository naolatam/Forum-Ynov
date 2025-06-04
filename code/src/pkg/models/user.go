package models

import (
	"time"

	"github.com/google/uuid"
)

// User represent a line in the users table
type User struct {
	ID         uuid.UUID
	Pseudo     string
	Email      string
	Password   string
	Bio        string
	Avatar     string
	CreatedAt  time.Time
	Role_ID    uuid.UUID
	Google_ID  *int
	Github_ID  *int
	Role       Role
	Session    Session
}
