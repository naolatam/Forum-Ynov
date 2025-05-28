package models

import (
	"github.com/google/uuid"
)

// Category représente une ligne dans la table category
type User struct {
	ID         uuid.UUID
	Pseudo     string
	Email      string
	Password   string
	Bio        string
	Avatar     string
	CreatedAt  string
	role_ID    uuid.UUID
	Role       Role
	session_ID string
	Session    Session
}
