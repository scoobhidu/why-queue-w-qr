package handlers

import (
	"database/sql"
	"fmt"
	"github.com/RashadAnsari/go-batch/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"reflect"
	"sync"
	"time"
	"why-queue-w-qr/models"
	"why-queue-w-qr/utils"
)

func AddExcusedAttendance(c *fiber.Ctx) error {
	fmt.Println(reflect.TypeOf(c.Params("class")))
	fmt.Println(reflect.TypeOf(c.Params("loginID")))
	fmt.Println(reflect.TypeOf(c.Params("enrolment_no")))

	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

var lati float64 = 77.06595815940048
var longi float64 = 28.71931254354033

var BatchMaster *batch.Batch[func()]
var attendanceDB *sql.DB

func ConnectAttendanceDB() {
	postgreConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("attendanceDBname"))

	log.Println(postgreConn)
	var err error

	attendanceDB, err = sql.Open("postgres", postgreConn)
	if err != nil {
		log.Fatalf("Error Connecting to the attendance DB: %s", err.Error())
	}
}

func CloseAttendanceDB() {
	err := attendanceDB.Close()
	if err != nil {
		log.Fatalf("Error closing your Database! %s", err)
	}
}

func MarkAttendance(c *fiber.Ctx) error {
	payload := new(models.AttendanceReqModel)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Wrong body passed %s", err),
		})
	}
	userTimestamp, _ := time.ParseDuration(payload.Timestamp)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go utils.CheckJWT(payload.JwtToken, userTimestamp.Milliseconds(), wg)
	wg.Add(1)
	go utils.GetDistanceFromLatLonInKm(payload.Lat, payload.Longi, lati, longi, wg)
	wg.Wait()

	if <-utils.JwtCheck && <-utils.DistanceCheck != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Either you are not on the location or you tried to scan an expired jwt",
		})
	}

	// else request validated and added to batchMaster
	BatchMaster.Input <- func() {
		fmt.Println("Starting Execution")
		query := fmt.Sprintf("update %s set %s = array_cat(%s, '{\"%s\"}') where enroll_no=%d;", payload.Class, payload.TeacherId, payload.TeacherId, payload.Timestamp, payload.EnrolmentNo)
		fmt.Println(query)
		_, err = attendanceDB.Query(query)
		switch {
		case err != nil:
			log.Fatalf("Err: %s", err)
		}
		fmt.Println("executed statement")
	}

	fmt.Printf("Payload: %s, %d, %s, %s, %s", payload.Class, payload.EnrolmentNo, payload.TeacherId, payload.JwtToken, payload.Timestamp)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func GetStudentAttendance(c *fiber.Ctx) error {
	studentClass := c.Params("class")
	teacher_id := c.Params("teacher_id")
	session_id := c.Params("session_id")
	enrolment_no := c.Params("enrolment_no")

	// session id == nil => directly go for querying
	// if both IDs are nil, status is not authenticated

	if session_id == "" && teacher_id == "" {
		// not authorized
	} else if session_id == "" {
		// direct query
	} else {
		// query with session_id check
	}

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
