package routes

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/1rvyn/takehome/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFile(c *fiber.Ctx) error {
	// get the session from the JWT token
	cookie := c.Cookies("jwt")

	if cookie == "" {
		return c.JSON(fiber.Map{
			"message": "Error cookie is empty",
		})
	}

	session, err := database.Redis.GetHMap(cookie)
	if err != nil {
		return err
	}

	if session == nil {
		return c.JSON(fiber.Map{
			"message": "Session error",
		})
	}

	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid file",
		})
	}

	// Create a unique file name
	fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))

	// Save the file to disk
	if err := c.SaveFile(file, fmt.Sprintf("./uploads/%s", fileName)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}

	// Update the user's file path
	userID, err := strconv.Atoi(session["user_id"])
	if err != nil {
		fmt.Println("Error converting string to int", userID)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server error",
		})
	}

	if user, ok := userMap[userID]; ok {
		user.FilePath = &fileName
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Return the user object
	return c.JSON(fiber.Map{
		"message": "Upload",
		"user":    userMap[userID],
	})
}
