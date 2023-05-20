package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/1rvyn/takehome/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// load environment variables
	var SecretKey = os.Getenv("SECRET_KEY")
	var SecretKey2 = os.Getenv("SECRET_KEY2")
	var SALT = os.Getenv("SALT")

	fmt.Println("Secret Key: ", SecretKey, "Salt: ", SALT, "secret2: ", SecretKey2)

	// Create the fiber app and setup routes
	app := fiber.New()

	setupRoutes(app)

	fmt.Println("Listening on port 8080")

	app.Listen(":8080")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", routes.Hello)

	app.Post("/login", routes.Login)
	app.Post("/register", routes.Register)

	app.Post("/create", routes.CreateUser)

	app.Get("/user/:id", routes.GetUser)
	app.Put("/update/:id", routes.UpdateUser)
	app.Delete("/user/:id", routes.DeleteUser)
}
