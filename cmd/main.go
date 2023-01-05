package main

import (
	goBatch "github.com/RashadAnsari/go-batch/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"time"
	"why-queue-w-qr/handlers"
	"why-queue-w-qr/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env data: %s", err.Error())
	}

	go handlers.ConnectStudentDB()
	go handlers.ConnectAttendanceDB()

	defer handlers.CloseStudentDB()
	defer handlers.CloseAttendanceDB()

	handlers.BatchMaster = goBatch.New[func()](
		goBatch.WithMaxWait(60*time.Second),
		goBatch.WithSize(100000),
	)

	go func() {
		for {
			output := <-handlers.BatchMaster.Output
			for _, f := range output {
				f()
			}
		}
	}()

	app := fiber.New()
	v1 := app.Group("/v1")

	v1.Post(utils.Login, handlers.LoginFunc)
	v1.Post(utils.MarkAttendance, handlers.MarkAttendance)
	v1.Post(utils.AddExcusedAttendance, handlers.AddExcusedAttendance)

	v1.Get(utils.GetStudentAttendance, handlers.GetStudentAttendance)
	v1.Get(utils.GetClassAttendance, handlers.GetClassAttendance)

	v1.Get(utils.GetAll, handlers.GetAll)

	log.Fatal(app.Listen("127.0.0.1:9010"))
}
