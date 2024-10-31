package models

import "time"

type Material struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	MediaURL    string     `json:"media_url"`
	AddedBy     int        `json:"added_by"`
	Status      StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:now()"`
}
