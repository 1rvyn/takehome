package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteUser(c *fiber.Ctx) error {
	// Parse the user ID from the request parameters
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	// Look up the user by ID in the userMap map
	_, ok := userMap[userID]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Remove the user from the userMap map
	delete(userMap, userID)

	// Return a response indicating success
	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
