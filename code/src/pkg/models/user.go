package models

import (
	"github.com/google/uuid"
)

// Category représente une ligne dans la table category
type User struct {
	ID        uuid.UUID
	Pseudo    string
	Email     string
	Password  string
	Bio       string
	Avatar    string
	CreatedAt string
	role_id   uuid.UUID
	Role      Role
}
