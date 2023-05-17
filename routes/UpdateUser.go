package routes

import "github.com/gofiber/fiber/v2"

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}
