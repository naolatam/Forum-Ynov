package models

import (
	"github.com/google/uuid"
)

// Reaction represent a line in the reactions table
type Reaction struct {
	ID         uuid.UUID
	Post_id    *uint32
	Post       Post
	Comment_id *uint32
	Comment    Comment
	user_id    uuid.UUID
	User       User
	Label      string
}
