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
	TimeAgo     string    // Human-readable time difference
	FromUser_ID uuid.UUID // User who created the notification
	User_ID     uuid.UUID
	Post_ID     uint32
	FromUser    User
	Post        Post
}
