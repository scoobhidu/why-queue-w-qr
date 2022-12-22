package utils

const Login string = "/login/:enrol_no/:session_id"
const GetAll string = "/login/all"
const MarkAttendance string = "/attendance/mark/:teacher_id/:jwt_token/:timestamp"
const GetStudentAttendance string = "/attendance/student/:enrolment_no/:session_id"
const GetClassAttendance string = "/attendance/class/:class"
const AddExcusedAttendance string = "/attendance/excused/:loginID/:class/:enrolment_no"
