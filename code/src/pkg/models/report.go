package models

import (
	"time"

	"github.com/google/uuid"
)

// Reaction represent a line in the reactions table
type Report struct {
	ID       uuid.UUID
	Post_id  uint32
	Post     Post
	User_id  uuid.UUID
	User     User
	ReportAt time.Time
}
