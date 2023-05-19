package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/1rvyn/takehome/routes"
)

// load env vars
var SecretKey = os.Getenv("SECRET_KEY")
var SALT = os.Getenv("SALT")

func main() {
	fmt.Println("Secret Key: ", SecretKey, "Salt: ", SALT)
	app := fiber.New()

	setupRoutes(app)

	fmt.Println("Listening on port 8080")

	app.Listen(":8080")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", routes.Hello)
	app.Post("/create", routes.CreateUser)
	app.Get("/user/:id", routes.GetUser)
	app.Put("/update/:id", routes.UpdateUser)
	app.Delete("/user/:id", routes.DeleteUser)
}
