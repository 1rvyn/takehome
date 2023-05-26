package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"

	"github.com/1rvyn/takehome/routes"

	"github.com/1rvyn/takehome/database"
)

func main() {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// load environment variables
	var SecretKey = os.Getenv("SECRET_KEY")
	var SecretKey2 = os.Getenv("SECRET_KEY2")
	var SALT = os.Getenv("SALT")

	// debug to see enviroment variables
	fmt.Println("Secret Key: ", SecretKey, "Salt: ", SALT, "secret2: ", SecretKey2)

	database.ConnectToRedis()

	// Create the fiber app and setup routes
	app := fiber.New()

	limiterMiddleware := limiter.New(limiter.Config{
		Max:        20,              // Maximum number of requests per window
		Expiration: 1 * time.Minute, // Time window for requests
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, http://127.0.0.1, localhost", // frontend url
		AllowHeaders:     "Origin, Content-Type, Accept, Set-Cookie, Cookie , Content-Type",
		AllowMethods:     "POST, OPTIONS, GET, PUT, DELETE, PREFLIGHT",
		AllowCredentials: true,
	}),
		limiterMiddleware, func(c *fiber.Ctx) error {
			if c.Method() == fiber.MethodPost {
				c.Set(fiber.HeaderContentLength, "4MB") // limit file size to 4mb (fiber has some kind of internal limit similar to uploadthingy from theo)
			}
			return c.Next()
		})

	// we setup all the routes and pass the app
	setupRoutes(app)

	fmt.Println("Listening on port 8080")

	app.Listen(":8080")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", routes.Hello)
	app.Post("/login", routes.Login)
	app.Post("/register", routes.Register)
	app.Get("/validate", routes.Validate) // used to keep a user logged in *we cant protect this as it needs to be accessed by non logged in users*

	// These are the protected routes which means each request is validated for a valid cookie + session
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
