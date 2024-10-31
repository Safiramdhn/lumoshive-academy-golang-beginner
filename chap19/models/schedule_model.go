package models

import "time"

type Schedule struct {
	ID        int        `json:"id"`
	Date      time.Time  `json:"date"`
	Time      time.Time  `json:"time"`
	AddedBy   int        `json:"added_by"`
	Status    StatusEnum `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
