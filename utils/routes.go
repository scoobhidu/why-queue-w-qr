package utils

const Login string = "/login/:enrol_no/:session_id"
const GetAll string = "/login/all"
const MarkAttendance string = "/attendance/mark"
const GetStudentAttendance string = "/attendance/student"
const GetClassAttendance string = "/attendance/class/:class/:teacher_id"
const AddExcusedAttendance string = "/attendance/excused/:class/:teacher_id/:enrolment_no"
