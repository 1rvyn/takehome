package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://127.0.0.1, localhost", // frontend url
		AllowHeaders:     "Origin, Content-Type, Accept, Set-Cookie, Cookie , Content-Type",
		AllowMethods:     "POST, OPTIONS, GET, PUT, DELETE, PREFLIGHT",
		AllowCredentials: true,
	}))

	// limit file uploads to 4MB
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodPost {
			c.Set(fiber.HeaderContentLength, "4MB")
		}
		return c.Next()
	})

	setupRoutes(app)

	fmt.Println("Listening on port 8080")

	app.Listen(":8080")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", routes.Hello)
	app.Post("/login", routes.Login)
	app.Post("/register", routes.Register)
	app.Get("/validate", routes.Validate)
	// app.Get("/users", routes.GetAllUsers)

	// app.Post("/create", routes.CreateUser)

	protected := app.Group("/")
	protected.Use(requireAuth())
	protected.Get("/users", routes.GetAllUsers)
	protected.Post("/create", routes.CreateUser)
	protected.Get("/user/:id", routes.GetUser)
	protected.Put("/update/:id", routes.UpdateUser)
	protected.Delete("/user/:id", routes.DeleteUser)
	protected.Post("/upload", routes.UploadFile)
	protected.Post("/logout", routes.Logout)

	app.Get("/session", routes.GetSession) // testing
}

func requireAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the request cookie
		cookie := c.Cookies("jwt")
		fmt.Println(cookie, " auth function cookie ")

		if cookie == "" {
			fmt.Println("33333")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access",
			})
		}

		// Get session from redis
		session, err := database.Redis.GetHMap(cookie)
		if err != nil {
			fmt.Println("22222")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access",
			})
		}

		if session == nil || session["user_id"] == "" {
			fmt.Println("44444")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized access",
			})
		}

		// Call the next handler
		return c.Next()
	}
}
