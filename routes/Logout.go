package routes

import (
	"fmt"

	"github.com/1rvyn/takehome/database"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	fmt.Println("logging out")
	// validate its a logged in user trying to logout
	cookie := c.Cookies("jwt")

	if cookie == "" {
		fmt.Println("no session cookie")
		return c.SendStatus(401)
	}

	// search for the session in redis
	session, err := database.Redis.GetHMap(cookie)
	if err != nil {
		fmt.Println("error getting session from redis")
		return c.SendStatus(401)
	}

	// fmt.Println("deleting the token", session["token"])
	// since there is a session, delete it from redis
	err = database.Redis.DeleteHMap(session["token"])
	if err != nil {
		fmt.Println("error deleting session from redis")
		return c.SendStatus(401)
	}

	// delete the session cookie
	c.ClearCookie("session")
	return c.SendStatus(200)
}
