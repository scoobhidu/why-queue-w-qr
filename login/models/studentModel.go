package models

import "time"

type Student struct {
	EnrolmentNo int       `json:"Enrolment No" field:"enrolmentno"`
	Name        string    `json:"Name" field:"studentname"`
	Dob         time.Time `json:"Date of Birth" field:"dob"`
	Branch      string    `json:"Branch" field:"branch"`
	Class       string    `json:"Class" field:"studentclass"`
	Subjects    []uint8   `json:"Subjects" field:"subjectcodes"`
	SessionID   string    `json:"Session ID" field:"sessionid"`
}
