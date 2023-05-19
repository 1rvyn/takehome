package routes

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/1rvyn/takehome/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {

	userID := c.Params("id")
	fmt.Println("user is ", userID)

	// parse the new user from the payload

	user := new(models.User)

	err := json.Unmarshal(c.Body(), user)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	// find the user in the users slice and update it

	update := false
	for i, u := range users {
		if strconv.Itoa(u.ID) == userID {
			users[i] = *user
			update = true
		}
	}

	// good practice
	if !update {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	fmt.Println("new users slice is ", users)

	return c.SendString("UpdateUser")
}
