package models

import "time"

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
	TotalOrders  int
}

type OrderSummary struct {
	Month       time.Time
	Id          int
	Name        string
	TotalOrders int
}
