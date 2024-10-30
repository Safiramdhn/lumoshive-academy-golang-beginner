package models

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Orders struct {
	Id           int
	CustomerId   int
	DriverId     int
	City         string
	District     string
	Neighborhood string
	StreetName   string
	OrderDate    time.Time
	OrderTime    time.Time
	OrderStatus  string
	Status       string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}
