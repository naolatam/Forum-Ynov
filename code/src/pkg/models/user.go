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
<<<<<<< HEAD
=======
	Google_ID  *int
	Github_ID  *int
>>>>>>> 3ceb4af2d34c11dd48114bfb7efa28a83563ae2e
	Role       Role
	session_ID string
	Session    Session
	Google_id  int
	Github_id  int
}
