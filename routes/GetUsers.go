package routes

import (
	"github.com/gofiber/fiber/v2"
)

// return all users
func GetAllUsers(c *fiber.Ctx) error {
	// Parse the user ID from the request parameters

	// Return the user object
	return c.JSON(userMap)
}
