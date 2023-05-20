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

	// Generate a new ID for the user
	newID := len(userMap) + 1 // We would typically use a UUID here in a DB where the DB assigns the ID
	newUser.ID = newID

	// Add the new user to the userMap map
	userMap[newID] = &newUser

	// Return a response indicating success
	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user":    newUser,
	})
}
