package models

type StatusEnum string

const (
	StatusActive  StatusEnum = "active"
	StatusDeleted StatusEnum = "deleted"
)

type AttendanceStatusEnum string

const (
	AttendanceOnTime AttendanceStatusEnum = "ontime"
	AttendanceLate   AttendanceStatusEnum = "late"
)

type AssignmentStatusEnum string

const (
	AssignmentNotStarted AssignmentStatusEnum = "not_started"
	AssignmentSubmitted  AssignmentStatusEnum = "submited"
	AssignmentMissing    AssignmentStatusEnum = "missing"
)
