package models

import "time"

type Student struct {
	ID        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	UserID    int        `json:"user_id"`
	Status    StatusEnum `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type StudentClass struct {
	ID        int        `json:"id"`
	StudentID int        `json:"student_id"`
	ClassID   int        `json:"class_id"`
	Status    StatusEnum `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
