package models

import (
	"time"

	"github.com/google/uuid"
)

// Category represents a line in the categories table
type RecentActivity struct {
	ID        uuid.UUID
	Action    string
	Details   string
	SubTitle  *string
	CreatedAt time.Time
	TimeAgo   string // This field is used to display the time in a human-readable format
	User_ID   uuid.UUID
	Post_ID   uint32
}
