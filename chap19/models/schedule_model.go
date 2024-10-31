package models

import "time"

type Schedule struct {
	ID        int        `json:"id"`
	Date      time.Time  `json:"date"`
	Time      time.Time  `json:"time"`
	AddedBy   int        `json:"added_by"`
	Status    StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now()"`
}
