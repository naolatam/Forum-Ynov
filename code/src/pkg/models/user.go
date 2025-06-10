package models

import (
	"time"

	"github.com/google/uuid"
)

// User represent a line in the users table
type User struct {
	ID        uuid.UUID
	Pseudo    string
	Email     string
	Password  string
	Bio       string
	Avatar    []byte
	CreatedAt time.Time
	Role_ID   uuid.UUID
	Google_ID *string
	Github_ID *int64
	Role      Role
	Session   Session
}
