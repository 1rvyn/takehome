package routes

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/1rvyn/takehome/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Define a map to store the session data

// func Login(c *fiber.Ctx) error {
// 	// we will need to parse the payload using the model.User struct
// 	// check the email exists in the users slice and then check the password

// 	// parse payload using the model.User struct
// 	user := new(models.User)
// 	err := json.Unmarshal(c.Body(), user)
// 	if err != nil {
// 		fmt.Println(err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request payload",
// 		})
// 	}

// 	fmt.Println("user: ", user)

// 	// check the email exists in the users map and then check the password
// 	if u, ok := userMap[user.ID]; ok {
// 		fmt.Println("userMap", userMap[user.ID], "the u: ", u)
// 		if u.Email == user.Email && u.Password == user.Password {
// 			fmt.Println("user is valid")
// 			// we create the cookie and set it here
// 			claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 				Issuer:    strconv.Itoa(user.ID),
// 				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 			})

// 			token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

// 			if err != nil {
// 				c.Status(fiber.StatusInternalServerError)
// 				return c.JSON(fiber.Map{
// 					"message": "could not create cookie",
// 				})
// 			}

// 			cookie := fiber.Cookie{
// 				Name:    "jwt",
// 				Value:   token,
// 				Expires: time.Now().Add(time.Hour * 24),
// 				MaxAge:  86400,
// 			}

// 			// this is where we set the cookie
// 			c.Cookie(&cookie)

// 			fmt.Println("we just made this cookie: ", cookie, "is is the datatype: ", reflect.TypeOf(cookie))

// 			// create a server side token to verify the JWT token

// 			SToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 				Issuer:    strconv.Itoa(user.ID),
// 				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
// 			}).SignedString([]byte(os.Getenv("SECRET_KEY2")))

// 			// they never get this token but we use it to validate their JWT token

// 			if err != nil {
// 				c.Status(fiber.StatusInternalServerError)
// 				return c.JSON(fiber.Map{
// 					"message": "failed to create session",
// 				})
// 			}

// 			session := make(map[string]interface{})
// 			session["user_id"] = user.ID
// 			session["email"] = user.Email
// 			session["token"] = SToken
// 			session["expires_at"] = time.Now().Add(time.Hour * 24).Unix()

// 			// we will create a K:V pair here to represent storing sessions in memory

// 			// Key = JWT token, Value = the session map we just made
// 			sessions[token] = session

// 			return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 				"message": "successful login",
// 				"user":    user,
// 			})
// 		}
// 	}

// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 		"message": "User not found",
// 	})

// }
func Login(c *fiber.Ctx) error {
	// Parse the request body into a struct
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.BodyParser(&loginData)
	if err != nil {
		return err
	}

	// Find the user by email
	for _, user := range userMap {
		if user.Email == loginData.Email {
			// Check if the password matches
			if user.Password == loginData.Password {
				// Return a response indicating success & create a session

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

				// we will create a K:V pair here to represent storing sessions

				// Key = JWT token, Value = the session map we just made
				err = database.Redis.PutHMap(token, session)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"message": "Failed to create session",
					})
				}

				return c.JSON(fiber.Map{
					"message": "Login successful",
					"user":    user,
				})
			} else {
				// Return a response indicating incorrect password
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Incorrect password",
				})
			}
		}
	}

	// Return a response indicating user not found
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "User not found",
	})
}
