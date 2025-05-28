package models

import (
	"github.com/google/uuid"
)

// Role represents a line in the role table
type Role struct {
	ID         uuid.UUID
	Name       string
	Permission int8
}
