package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Id          int
	Email       string
	Password    string
	User_id     int
	Login_time  sql.NullTime
	Logout_time sql.NullTime
	Status      string
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
}
