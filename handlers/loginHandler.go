package handlers

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"strconv"
	"why-queue-w-qr/models"
)

var studentsDB *sql.DB

func ConnectStudentDB() {
	postgreConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("studentDBname"))

	log.Println(postgreConn)
	var err error

	studentsDB, err = sql.Open("postgres", postgreConn)
	if err != nil {
		log.Fatalf("Error Connecting to the students DB: %s", err.Error())
	}
}

func CloseStudentDB() {
	err := studentsDB.Close()
	if err != nil {
		log.Fatalf("Error closing your Database! %s", err)
	}
}

func GetAll(c *fiber.Ctx) error {
	println("Querying started")
	rows, err := studentsDB.Query("select * from students")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("Error %s", err),
		})
	}
	var students []models.Student

	for rows.Next() {
		var stu models.Student
		if err := rows.Scan(&stu.EnrolmentNo, &stu.Name, &stu.Dob, &stu.Branch, &stu.Class, &stu.Subjects, &stu.SessionID, &stu.YearOfPassing); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": fmt.Sprintf("Error2 : %s", err),
			})
		}
		students = append(students, stu)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "Working",
		"students": students,
	})
}

func LoginFunc(c *fiber.Ctx) error {
	enrol_no, err := strconv.Atoi(c.Params("enrol_no"))
	sessionId := c.Params("session_id")

	println("Login Querying started")

	var dbSessionId sql.NullString
	err = studentsDB.QueryRow(fmt.Sprintf("select session_id from students where enroll_no = %d", enrol_no)).Scan(&dbSessionId)
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
		_, err := studentsDB.Exec(fmt.Sprintf("UPDATE students SET session_id = '%s' WHERE enroll_no = %d", sessionId, enrol_no))
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
