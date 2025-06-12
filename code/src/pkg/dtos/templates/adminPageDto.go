package dtos

import "Forum-back/pkg/models"

type AdminPageDto struct {
	Header        HeaderDto           `json:"header"`
	AllUsers      []models.User       `json:"all_users"`
	AllPost       *[]*models.Post     `json:"all_post"`
	AllCategories *[]*models.Category `json:"all_category"`
}
