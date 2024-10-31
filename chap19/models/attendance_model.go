package models

import "time"

type MentorAttendance struct {
	ID               int                  `json:"id"`
	MentorID         int                  `json:"mentor_id"`
	ClassID          int                  `json:"class_id"`
	DateTime         time.Time            `json:"datetime"`
	AttendanceStatus AttendanceStatusEnum `json:"attendance_status"`
	Status           StatusEnum           `json:"status" gorm:"default:'active'"`
	CreatedAt        time.Time            `json:"created_at" gorm:"default:now()"`
	UpdatedAt        time.Time            `json:"updated_at" gorm:"default:now()"`
}

type StudentAttendance struct {
	ID               int                  `json:"id"`
	StudentID        int                  `json:"student_id"`
	StudentClassID   int                  `json:"student_class_id"`
	DateTime         time.Time            `json:"datetime"`
	AttendanceStatus AttendanceStatusEnum `json:"attendance_status"`
	Status           StatusEnum           `json:"status" gorm:"default:'active'"`
	CreatedAt        time.Time            `json:"created_at" gorm:"default:now()"`
	UpdatedAt        time.Time            `json:"updated_at" gorm:"default:now()"`
}
