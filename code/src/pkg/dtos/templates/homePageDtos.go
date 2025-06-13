package dtos

import "Forum-back/pkg/models"

type HomePageDto struct {
	Header           HeaderDto
	LastPosts        []*models.Post
	UserCount        int
	PostCount        int
	ActiveUsersCount int
}
