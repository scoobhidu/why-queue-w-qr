package attendanceHandlers

import (
	"database/sql"
	"fmt"
	"github.com/RashadAnsari/go-batch/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"log"
	"reflect"
	"sync"
	"why-queue-w-qr/attendance/models/AttendanceReqBodyModel"
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
var longi = 28.71931254354033

var BatchMaster *batch.Batch[func()]
var AttendanceDB *sql.DB

func MarkAttendance(c *fiber.Ctx) error {
	payload := new(AttendanceReqBodyModel.AttendanceReqModel)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fasthttp.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Wrong body passed %s", err),
		})
	}

	// TODO check with mongoDB for JWT and timestamp
	// TODO + location
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go utils.CheckJWT(payload.JwtToken, payload.Timestamp, wg)
	wg.Add(1)
	go utils.GetDistanceFromLatLonInKm()

	wg.Wait()

	BatchMaster.Input <- func() {
		fmt.Println("Starting Execution")
		query := fmt.Sprintf("update %s set %s = array_cat(%s, '{\"%s\"}') where enroll_no=%d;", payload.Class, payload.TeacherId, payload.TeacherId, payload.Timestamp, payload.EnrolmentNo)
		fmt.Println(query)
		_, err = AttendanceDB.Query(query)
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
