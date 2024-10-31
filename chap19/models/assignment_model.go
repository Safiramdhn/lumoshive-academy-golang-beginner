package models

import "time"

type Assignment struct {
	ID          int        `json:"id"`
	ClassID     int        `json:"class_id"`
	Deadline    time.Time  `json:"deadline"`
	Description string     `json:"description"`
	Title       string     `json:"title"`
	Status      StatusEnum `json:"status" gorm:"default:'active'"`
	CreatedAt   time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"default:now()"`
}

type StudentAssignment struct {
	ID               int                  `json:"id"`
	AssignmentID     int                  `json:"assignment_id"`
	StudentID        int                  `json:"student_id"`
	SubmitDate       *time.Time           `json:"submit_date"`
	AssignmentStatus AssignmentStatusEnum `json:"assignment_status" gorm:"default:'not_started'"`
	Score            *float64             `json:"score"`
	Status           StatusEnum           `json:"status" gorm:"default:'active'"`
	CreatedAt        time.Time            `json:"created_at" gorm:"default:now()"`
	UpdatedAt        time.Time            `json:"updated_at" gorm:"default:now()"`
}
