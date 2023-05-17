package routes

import "github.com/gofiber/fiber/v2"

// GetUser is a function that returns a user
func GetUser(c *fiber.Ctx) error {
	return c.SendString("Get User")
}
