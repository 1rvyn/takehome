package routes

import "github.com/gofiber/fiber/v2"

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
