package models

import "time"

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Status    StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:now()"`
}
