package models

import (
	"database/sql"
)

type User struct {
	Id          int
	Email       string
	Password    string
	User_id     int
	Login_time  sql.NullTime
	Logout_time sql.NullTime
}
