package dtos

import "html/template"

type RecentActivityDto struct {
	Action      string `json:"activity_title"`
	Description template.HTML
	TimeAgo     string `json:"time_ago"`
}
