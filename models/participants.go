package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	RoomID    uuid.UUID `json:"RoomID" form:"RoomID"`
	CreateAt  time.Time
	CreateBy  string
	UpdateAt  sql.NullTime
	UpdateBy  sql.NullString
	DeleteAt  sql.NullTime
	DeleteBy  sql.NullString
	RecStatus int
}
