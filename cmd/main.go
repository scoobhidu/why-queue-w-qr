package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"why-queue-w-qr/login/handlers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env data: %s", err.Error())
	}
}

func main() {
	postgreConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"))
	var err error

	handlers.StudentsDB, err = sql.Open("postgres", postgreConn)
	if err != nil {
		log.Fatalf("Error Connecting to the students DB: %s", err.Error())
	}

	defer func(StudentsDB *sql.DB) {
		err := StudentsDB.Close()
		if err != nil {
			log.Fatalf("Error closing your Database! %s", err)
		}
	}(handlers.StudentsDB)

	app := fiber.New()
	v1 := app.Group("/v1")

	const login string = "/login/:enrol_no/:session_id"
	const loginAll string = "/login/all"
	const markAttendance string = "/attendance/mark/:teacher_id/:jwt_token/:timestamp"
	const getStudentAttendance string = "/attendance/student/:enrolment_no/:session_id"
	const getClassAttendance string = "/attendance/class/:class"
	const addExcusedAttendance string = "/attendance/excused/:class/:enrolment_no"

	v1.Post(login, handlers.LoginFunc)
	v1.Post(markAttendance, handlers.MarkAttendance)
	v1.Post(addExcusedAttendance, handlers.AddExcusedAttendance)
	v1.Get(loginAll, handlers.GetAll)
	v1.Get(getStudentAttendance, handlers.GetStudentAttendance)
	v1.Get(getClassAttendance, handlers.GetClassAttendance)

	log.Fatal(app.Listen("127.0.0.1:9010"))
}
