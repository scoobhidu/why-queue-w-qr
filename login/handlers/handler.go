package handlers

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strconv"
	"why-queue-w-qr/login/models"
)

var StudentsDB *sql.DB

func GetAll(c *fiber.Ctx) error {
	println("Querying started")
	rows, err := StudentsDB.Query("select * from student")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Error %s", err),
		})
	}
	var students []models.Student

	for rows.Next() {
		var stu models.Student
		if err := rows.Scan(&stu.EnrolmentNo, &stu.Name, &stu.Dob, &stu.Branch, &stu.Class, &stu.Subjects, &stu.SessionID, &stu.YearOfPassing); err != nil {
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

func LoginFunc(c *fiber.Ctx) error {
	enrol_no, err := strconv.Atoi(c.Params("enrol_no"))
	sessionId := c.Params("session_id")

	println("Login Querying started")

	var dbSessionId sql.NullString
	err = StudentsDB.QueryRow(fmt.Sprintf("select session_id from student where enroll_no = %d", enrol_no)).Scan(&dbSessionId)
	switch {
	case err == sql.ErrNoRows:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong credentials",
		})

	case err != nil:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error %s", err),
		})

	case dbSessionId.Valid == false:
		_, err := StudentsDB.Exec(fmt.Sprintf("UPDATE student SET session_id = '%s' WHERE enroll_no = %d", sessionId, enrol_no))
		fmt.Println("Updating")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": fmt.Sprintf("Error %s", err),
			})
		}

	case dbSessionId.String != sessionId:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please don't try to login for your friend",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func AddExcusedAttendance(c *fiber.Ctx) error {
	fmt.Println(reflect.TypeOf(c.Params("class")))
	fmt.Println(reflect.TypeOf(c.Params("loginID")))
	fmt.Println(reflect.TypeOf(c.Params("enrolment_no")))

	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

func MarkAttendance(c *fiber.Ctx) error {
	fmt.Println(reflect.TypeOf(c.Params("class")))
	fmt.Println(reflect.TypeOf(c.Params("loginID")))
	fmt.Println(reflect.TypeOf(c.Params("enrolment_no")))
	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

func GetStudentAttendance(c *fiber.Ctx) error {
	fmt.Println(reflect.TypeOf(c.Params("class")))
	fmt.Println(reflect.TypeOf(c.Params("loginID")))
	fmt.Println(reflect.TypeOf(c.Params("enrolment_no")))
	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

func GetClassAttendance(c *fiber.Ctx) error {
	fmt.Println(reflect.TypeOf(c.Params("class")))
	fmt.Println(reflect.TypeOf(c.Params("loginID")))
	fmt.Println(reflect.TypeOf(c.Params("enrolment_no")))
	return c.JSON(fiber.Map{
		"message": "Working",
	})
}
