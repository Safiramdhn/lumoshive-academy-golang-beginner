package models

import "time"

type Class struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AddedBy     int        `json:"added_by"`
	MentorID    int        `json:"mentor_id"`
	ScheduleID  int        `json:"schedule_id"`
	MaterialID  int        `json:"material_id"`
	Status      StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:now()"`
}

type Leaderboard struct {
	ID              int        `json:"id"`
	StudentID       int        `json:"student_id"`
	ClassID         int        `json:"class_id"`
	TotalScore      float64    `json:"total_score"`
	TotalAttendance int        `json:"total_attendance"`
	Status          StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt       time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"default:now()"`
}
