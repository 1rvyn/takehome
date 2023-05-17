package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/1rvyn/takehome/routes"
)

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Post("/create", routes.CreateUser)
	app.Get("/user/:id", routes.GetUser)
	app.Put("/user/:id", routes.UpdateUser)
	app.Delete("/user/:id", routes.DeleteUser)
}
