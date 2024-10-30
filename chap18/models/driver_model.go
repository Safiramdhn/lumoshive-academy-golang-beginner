package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Driver struct {
	Id        int
	FirstName string
	LastName  string
	UserId    int
	Status    string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
