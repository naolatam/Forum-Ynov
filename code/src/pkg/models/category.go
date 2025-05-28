package models

import (
	"github.com/google/uuid"
)

// Category représente une ligne dans la table category
type Category struct {
	ID   uuid.UUID
	Name string
}
