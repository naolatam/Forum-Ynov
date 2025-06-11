package dtos

import (
	"Forum-back/pkg/models"

	"github.com/google/uuid"
)

type PostPageDto struct {
	Header       HeaderDto `json:"header"`
	Post         models.Post
	Comments     []*CommentDto
	Like         int
	Dislike      int
	ActualUserId uuid.UUID
	UserReaction *models.Reaction
}

type CommentDto struct {
	Content      string
	ID           uint32
	User         models.User
	Date         string
	Like         int
	Dislike      int
	UserReaction *models.Reaction
}
