package models

import (
	"github.com/google/uuid"
)

// Reaction represent a line in the user reaction
type Reaction struct {
	ID         int64
	post_id    int64
	Post       Post
	comment_id int64
	Comment    Comment
	user_id    uuid.UUID
	User       User
	Label      string
}
