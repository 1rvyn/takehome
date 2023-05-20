package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
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

	// Parse the request body into the user object
	err = c.BodyParser(user)
	if err != nil {
		return err
	}

	// Return a response indicating success
	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}
