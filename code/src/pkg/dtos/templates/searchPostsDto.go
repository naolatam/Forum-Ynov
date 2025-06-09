package dtos

import (
	"Forum-back/pkg/models"

	"github.com/google/uuid"
)

type SearchPostsDto struct {
	Header         HeaderDto
	Posts          []*models.Post
	Categories     []*models.Category
	SearchTerm     string
	SearchCategory uuid.UUID
}
