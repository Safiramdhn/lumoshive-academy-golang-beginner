package models

import "time"

type Announcement struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AddedBy     int        `json:"added_by"`
	ClassID     int        `json:"class_id"`
	Status      StatusEnum `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
