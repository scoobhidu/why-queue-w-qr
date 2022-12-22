package attendanceHandlers

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"reflect"
	"why-queue-w-qr/attendance/models/AttendanceReqBodyModel"
)

var MongoCTX context.Context
var MongoClient *mongo.Client

func AddExcusedAttendance(c *fiber.Ctx) error {
	fmt.Println(reflect.TypeOf(c.Params("class")))
	fmt.Println(reflect.TypeOf(c.Params("loginID")))
	fmt.Println(reflect.TypeOf(c.Params("enrolment_no")))

	return c.JSON(fiber.Map{
		"message": "Working",
	})
}

func MarkAttendance(c *fiber.Ctx) error {
	payload := new(AttendanceReqBodyModel.AttendanceReqModel)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Wrong body passed %s", err),
		})
	}

	hackDB := MongoClient.Database(os.Getenv("mongoDBname"))
	jwtCollection := hackDB.Collection("qraccesses")

	fmt.Println(c.Query("token"))
	fmt.Printf("Payload: %s, %d, %s, %s, %d", payload.Class, payload.EnrolmentNo, payload.TeacherId, payload.JwtToken, payload.Timestamp)

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
