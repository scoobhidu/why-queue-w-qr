package handlers

import "github.com/gofiber/fiber/v2"

func LoginFunc(c *fiber.Ctx) error {
	enrol_id := c.Params("enrol_no")
	session_id := c.Params("session_id")

	// ? go routine TODO: Apply a bloom filter for checking the presence of session ID
	// ? go routine TODO: Apply a bloom filter for checking the presence of the enrolment no.

	// * login check TODO: if both are positive acc to bloom filter we check in the database
	// * signup check TODO: if the session id is not present for the given enrolment no. ! else 402

	// TODO: if the session id matches for the given enrolment no. allow login ! else 402

	return nil
}
