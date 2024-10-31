package models

import "time"

type Admin struct {
	ID        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	UserID    int        `json:"user_id"`
	Status    StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now()"`
}
