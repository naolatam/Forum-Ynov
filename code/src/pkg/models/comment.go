package models

import (
	"time"

	"github.com/google/uuid"
)

// Comment represent a line in the comments table
type Comment struct {
	ID        uint32
	Content   string
	CreatedAt time.Time
	Post_id   uint32
	Post      Post
	User_ID   uuid.UUID
	User      User
}
