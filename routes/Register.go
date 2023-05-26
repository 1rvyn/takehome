package routes

import (
	"github.com/1rvyn/takehome/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// Parse the request body into a new User struct
	var newUser models.User
	err := c.BodyParser(&newUser)
	if err != nil {
		return err
	}

	// Check if the email already exists in userMap
	for _, user := range userMap {
		if user.Email == newUser.Email {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "email already in use",
			})
		}
	}

	// Generate a new ID for the user
	newID := len(userMap) + 1
	newUser.ID = newID

	// Add the new user to the userMap map
	userMap[newID] = &newUser

	// Return a response indicating success
	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user":    newUser,
	})
}
