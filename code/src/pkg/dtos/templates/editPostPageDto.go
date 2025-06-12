package dtos

import (
	"Forum-back/pkg/models"
)

type EditPostPageDto struct {
	Header       HeaderDto `json:"header"`
	Post         models.Post
	User         models.User
	PostedDate   string
	Categories   []*models.Category
	UserReaction *models.Reaction
	Like         int
	Dislike      int
	IsNew        bool
}
