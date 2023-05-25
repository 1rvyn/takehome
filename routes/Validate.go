package routes

import (
	"fmt"

	"github.com/1rvyn/takehome/database"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx) error {
	// validate the user
	// 1. check if the user exists via redis

	cookie := c.Cookies("jwt")

	session, err := database.Redis.GetHMap(cookie)

	if err != nil {
		fmt.Println("error: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": false,
		})
	}

	if session == nil || session["user_id"] == "" {
		fmt.Println("session: ", session)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": false,
		})
	}

	if session != nil {
		fmt.Println("session: ", session)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": true,
		})
	}

	return c.JSON(fiber.Map{
		"message": false,
	})
}
