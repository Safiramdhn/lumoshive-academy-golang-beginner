package models

import "time"

type Announcement struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AddedBy     int        `json:"added_by"`
	ClassID     int        `json:"class_id"`
	Status      StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:now()"`
}
