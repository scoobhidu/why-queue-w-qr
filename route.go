package main

import (
	"github.com/gofiber/fiber/v2"
	"why-queue-w-qr/login/handlers"
)

const login string = "/login/:enrol_no/:session_id"
const markAttendance string = "/attendance/mark/"
const getStudentAttendance string = "/attendance/student"
const getClassAttendance string = "/attendance/class"

func setupRoutes(app fiber.Router) {
	app.Post(login, handlers.LoginFunc)

	// TODO add the required handler functions
	app.Post(markAttendance)
	app.Get(getStudentAttendance)
	app.Get(getClassAttendance)
}
