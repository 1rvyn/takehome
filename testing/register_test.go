package testing

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/1rvyn/takehome/models"
	"github.com/1rvyn/takehome/routes"
	"github.com/gofiber/fiber/v2"
)

func TestRegister(t *testing.T) {
	app := fiber.New()

	app.Post("/register", routes.Register)

	userMap := make(map[int]*models.User)

	testCases := []struct {
		name string
		user *models.User
		want string
	}{
		{
			name: "Successfully register new user",
			user: &models.User{
				Email:     "testuser1@example.com",
				FirstName: "Test User1",
				Password:  "testpassword",
			},
			want: "User created successfully",
		},
		{
			name: "Fail to register with existing email",
			user: &models.User{
				Email:     "testuser1@example.com",
				FirstName: "Test User2",
			},
			want: "email already in use",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.user)
			req := httptest.NewRequest("POST", "/register", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			if err != nil {
				t.Errorf("Failed to execute request: %v", err)
			}

			var result map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				t.Errorf("Failed to decode response: %v", err)
			}

			if result["message"] != tc.want {
				t.Errorf("Expected message %v, but got %v", tc.want, result["message"])
			}

			if result["message"] == "User created successfully" {
				userData := result["user"].(map[string]interface{})
				userMap[int(userData["ID"].(float64))] = &models.User{
					ID:    int(userData["ID"].(float64)),
					Email: userData["Email"].(string),
				}
			}
		})
	}
}
