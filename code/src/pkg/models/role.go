package models

import (
	"github.com/google/uuid"
)

// Category représente une ligne dans la table category
type Role struct {
	ID         uuid.UUID
	Name       string
	Permission int8
}
