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
	Status      StatusEnum `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
