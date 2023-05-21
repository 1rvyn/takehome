package routes

import (
	"fmt"

	"github.com/1rvyn/takehome/database"
	"github.com/gofiber/fiber/v2"
)

func GetSession(c *fiber.Ctx) error {

	// get the session from the JWT token

	// parse the JWT
	cookie := c.Cookies("jwt")

	fmt.Println("Cookie is: ", cookie)

	// search redis for the cookie to find the session

	session, err := database.Redis.GetHMap(cookie)
	if err != nil {
		return err
	}

	fmt.Println("Redis found this session: ", session)

	if session == nil {
		fmt.Println("there was no session found however there was a cookie (potentially expired / malicious)")
		return c.JSON(fiber.Map{
			"message": "Error",
		})
	}

	// Return the user object
	return c.JSON(fiber.Map{
		"message": "Cookie found",
	})

}
