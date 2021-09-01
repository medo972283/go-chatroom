package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID
	Content   string `form:"Content" json:"Content"`
	UserID    uuid.UUID
	RoomID    uuid.UUID `form:"RoomID" json:"RoomID"`
	CreateAt  time.Time
	CreateBy  string
	UpdateAt  sql.NullTime
	UpdateBy  sql.NullString
	DeleteAt  sql.NullTime
	DeleteBy  sql.NullString
	RecStatus int
}
