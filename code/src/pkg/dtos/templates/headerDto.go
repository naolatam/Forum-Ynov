package dtos

import "Forum-back/pkg/models"

type HeaderDto struct {
	IsConnected   bool
	IsAdmin       bool
	IsModerator   bool
	Notifications []*models.Notification
	PageName      string
}
