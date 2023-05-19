package routes

import (
	"encoding/json"
	"fmt"

	"github.com/1rvyn/takehome/models"
	"github.com/gofiber/fiber/v2"
)

var users []models.User

func CreateUser(c *fiber.Ctx) error {
	// parse payload using the model.User struct
	user := new(models.User)
	err := json.Unmarshal(c.Body(), user)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	// create new user
	users = append(users, *user)

	// return success response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}
