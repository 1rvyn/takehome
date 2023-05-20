package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetUser is a function that returns a user
func GetUser(c *fiber.Ctx) error {
	// Parse the user ID from the request parameters
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// Look up the user by ID in the userMap map
	user, ok := userMap[userID]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Return the user object
	return c.JSON(user)
}
