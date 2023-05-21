package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/1rvyn/takehome/routes"

	"github.com/1rvyn/takehome/database"
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

	database.ConnectToRedis()

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

	// app.Post("/create", routes.CreateUser)

	protected := app.Group("/")
	protected.Use(requireAuth())
	protected.Post("/create", routes.CreateUser)
	protected.Get("/user/:id", routes.GetUser)
	protected.Put("/update/:id", routes.UpdateUser)
	protected.Delete("/user/:id", routes.DeleteUser)
	protected.Get("/users", routes.GetAllUsers)
	protected.Post("/upload", routes.UploadFile)

	app.Get("/session", routes.GetSession) // testing
}

func requireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the request cookie
		cookie := c.Cookies("jwt")

		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access",
			})
		}

		// Get session from redis
		session, err := database.Redis.GetHMap(cookie)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access",
			})
		}

		if session == nil || session["user_id"] == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access",
			})
		}

		// Call the next handler
		return c.Next()
	}
}
