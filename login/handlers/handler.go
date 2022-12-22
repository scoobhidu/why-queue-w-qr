package handlers

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"why-queue-w-qr/login/models"
)

var StudentsDB *sql.DB

func LoginFunc(c *fiber.Ctx) error {
	//enrol_id := c.Params("enrol_no")
	//session_id := c.Params("session_id")

	// * login check TODO: if both are positive acc to bloom filter we check in the database
	// * signup check TODO: if the session id is not present for the given enrolment no. ! else 402

	// TODO: if the session id matches for the given enrolment no. allow login ! else 402

	return nil
}

func GetAll(c *fiber.Ctx) error {
	println("Querying started")
	rows, err := StudentsDB.Query("select * from students")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Error %s", err),
		})
	}
	var students []models.Student

	for rows.Next() {
		var stu models.Student
		if err := rows.Scan(&stu.EnrolmentNo, &stu.Name, &stu.Dob, &stu.Branch, &stu.Class, &stu.Subjects, &stu.SessionID); err != nil {
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("Error2 : %s", err),
			})
		}
		students = append(students, stu)
	}

	return c.JSON(fiber.Map{
		"message":  "Working",
		"students": students,
	})
}

func AddExcusedAttendance(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

func MarkAttendance(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

func GetStudentAttendance(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

func GetClassAttendance(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Working",
	})
}
