package models

import (
	"time"

	"github.com/google/uuid"
)

// Post represent a line in the posts table
type Post struct {
	ID        uint32
	Title     string
	Content   string
	validated bool
	CreatedAt time.Time
	user_ID   uuid.UUID
	User      User
}
