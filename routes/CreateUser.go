package routes

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return c.SendString("Create User")
}
