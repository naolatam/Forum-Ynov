package dtos

import "html/template"

type ProfilPageDto struct {
	Header         HeaderDto            `json:"header"`
	Username       string               `json:"username"`
	Email          string               `json:"email"`
	JoinedAt       string               `json:"joined_at"`
	Avatar         template.URL         `json:"avatar"`
	Bio            string               `json:"bio"`
	PostsCount     int                  `json:"posts_count"`
	CommentsCount  int                  `json:"comments_count"`
	RecentActivity []*RecentActivityDto `json:"recent_activity"`
	Error          ProfilePageErrorDto  `json:"error_string,omitempty"`
}

type ProfilePageErrorDto struct {
	ErrorTitle   string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}
