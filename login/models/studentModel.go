package models

import "time"

type Student struct {
	EnrolmentNo int64     `json:"Enrolment No"`
	Name        string    `json:"Name"`
	Dob         time.Time `json:"Date of Birth"`
	Branch      string    `json:"Branch"`
	Class       string    `json:"Class"`
	Subjects    []string  `json:"Subjects"`
	SessionID   string    `json:"Session ID"`
}
