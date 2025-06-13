package dtos

import "Forum-back/pkg/models"

type AdminPageDto struct {
	Header             HeaderDto           `json:"header"`
	AllUsers           []*models.User      `json:"all_users"`
	WaitingPosts       *[]*models.Post     `json:"waiting_posts"`
	Reports            *[]*models.Report   `json:"waiting_posts"`
	AllCategories      *[]*models.Category `json:"all_category"`
	UserManagement     bool                `json:"user_management"`
	ContentManagement  bool                `json:"content_management"`
	CategoryManagement bool                `json:"category_management"`
}
