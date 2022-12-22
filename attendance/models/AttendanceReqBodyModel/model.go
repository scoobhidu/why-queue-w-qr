package AttendanceReqBodyModel

type AttendanceReqModel struct {
	Class       string `json:"class" form:"class"`
	TeacherId   string `json:"teacher_id" form:"teacher_id"`
	EnrolmentNo int64  `json:"enrolment_no" form:"enrolment_no"`
	JwtToken    string `json:"token" form:"token"`
	Timestamp   int    `json:"time" form:"time"`
}
