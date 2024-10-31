package models

import "time"

type Mentor struct {
	ID        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	UserID    int        `json:"user_id"`
	AddedBy   int        `json:"added_by"`
	Status    StatusEnum `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
