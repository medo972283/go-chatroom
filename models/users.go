package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Account   string `form:"account" json:"account"`
	Password  string `form:"password" json:"password"`
	Nickname  string `form:"nickname" json:"nickname"`
	Email     string `form:"email" json:"email"`
	CreateAt  time.Time
	CreateBy  string
	UpdateAt  sql.NullTime
	UpdateBy  sql.NullString
	DeleteAt  sql.NullTime
	DeleteBy  sql.NullString
	RecStatus int
}
