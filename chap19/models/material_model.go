package models

import "time"

type Material struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	MediaURL    string     `json:"media_url"`
	AddedBy     int        `json:"added_by"`
	Status      StatusEnum `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
