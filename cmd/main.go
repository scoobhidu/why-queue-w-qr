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
	"why-queue-w-qr/utils"
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

	v1.Post(utils.Login, handlers.LoginFunc)
	v1.Post(utils.MarkAttendance, handlers.MarkAttendance)
	v1.Get(utils.AddExcusedAttendance, handlers.AddExcusedAttendance) // TODO change to POST
	v1.Get(utils.GetAll, handlers.GetAll)
	v1.Get(utils.GetStudentAttendance, handlers.GetStudentAttendance)
	v1.Get(utils.GetClassAttendance, handlers.GetClassAttendance)

	log.Fatal(app.Listen("127.0.0.1:9010"))
}
