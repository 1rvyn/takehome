package routes

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/1rvyn/takehome/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Define a map to store the session data
var sessions = make(map[string]map[string]interface{})

func Login(c *fiber.Ctx) error {
	// we will need to parse the payload using the model.User struct
	// check the email exists in the users slice and then check the password

	// parse payload using the model.User struct
	user := new(models.User)
	err := json.Unmarshal(c.Body(), user)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request payload",
		})
	}

	// check the email exists in the users map and then check the password
	if u, ok := userMap[user.ID]; ok {
		if u.Email == user.Email && u.Password == user.Password {
			// we create the cookie and set it here
			claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
				Issuer:    strconv.Itoa(user.ID),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			})

			token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(fiber.Map{
					"message": "could not create cookie",
				})
			}

			cookie := fiber.Cookie{
				Name:    "jwt",
				Value:   token,
				Expires: time.Now().Add(time.Hour * 24),
				MaxAge:  86400,
			}

			// this is where we set the cookie
			c.Cookie(&cookie)

			fmt.Println("we just made this cookie: ", cookie, "is is the datatype: ", reflect.TypeOf(cookie))

			// create a server side token to verify the JWT token

			SToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
				Issuer:    strconv.Itoa(user.ID),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			}).SignedString([]byte(os.Getenv("SECRET_KEY2")))

			// they never get this token but we use it to validate their JWT token

			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(fiber.Map{
					"message": "failed to create session",
				})
			}

			session := make(map[string]interface{})
			session["user_id"] = user.ID
			session["email"] = user.Email
			session["token"] = SToken
			session["expires_at"] = time.Now().Add(time.Hour * 24).Unix()

			// we will create a K:V pair here to represent storing sessions in memory

			// Key = JWT token, Value = the session map we just made
			sessions[token] = session

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "successful login",
				"user":    user,
			})
		}
	}

	fmt.Println("new session", sessions)

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "User not found",
	})

}
