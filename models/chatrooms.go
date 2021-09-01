package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Chatroom struct {
	ID        uuid.UUID
	Name      string `form:"chatroomName" json:"chatroomName"`
	CreateAt  time.Time
	CreateBy  string
	UpdateAt  sql.NullTime
	UpdateBy  sql.NullString
	DeleteAt  sql.NullTime
	DeleteBy  sql.NullString
	RecStatus int
}
