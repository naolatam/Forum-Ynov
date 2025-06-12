package models

import (
	"time"

	"github.com/google/uuid"
)

// Category represents a line in the categories table
type Notification struct {
	ID          int64
	Title       string
	Description string
	CreatedAt   time.Time
	TimeAgo     string // Human-readable time difference
	User_ID     uuid.UUID
}
