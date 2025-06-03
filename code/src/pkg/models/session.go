package models

import (
	"time"

	"github.com/google/uuid"
)

// Session represent a line in the sessions table
type Session struct {
	ID       uuid.UUID
	ExpireAt time.Time
	User_ID  uuid.UUID
	Expired  bool
}
