package models

import (
	"github.com/google/uuid"
)

// Category represents a line in the categorys table
type Category struct {
	ID   uuid.UUID
	Name string
}
