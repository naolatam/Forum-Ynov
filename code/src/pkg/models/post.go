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
	Validated bool
	CreatedAt time.Time
	User_ID   uuid.UUID
	User      User
}
