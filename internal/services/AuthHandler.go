package services

import (
	"cookies_session_demo/internal/dto"
	"cookies_session_demo/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

var Users = []models.Users{
	{
		Username:  "admin",
		Password:  "123",
		Sid:       "SID123456",
		Ipaddress: "127.0.0.1",
	},
	{
		Username:  "thanh",
		Password:  "456",
		Sid:       "SID789012",
		Ipaddress: "192.168.1.10",
	},
}

type AuthHandler struct{}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input dto.LoginRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	for _, u := range Users {
		if u.Username == input.Username && u.Password == input.Password {
			cookie := c.Cookies("my_cookie")
			if cookie == "" {
				return c.SendString("No cookie found")
			}
			resp := fiber.Map{
				"username":  u.Username,
				"sid":       u.Sid,
				"ipaddress": u.Ipaddress,
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user login successfully", "user": resp})
		}

	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "wrong username or password"})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input dto.RegisterRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	for _, u := range Users {
		if u.Username == input.Username {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "username already exists"})
		}
	}

	newUser := models.Users{
		Username:  input.Username,
		Password:  input.Password,
		Sid:       "SID" + time.Now().Format("150405"),
		Ipaddress: c.IP(),
	}

	Users = append(Users, newUser)
	c.Cookie(&fiber.Cookie{
		Name:     "my_cookie",
		Value:    newUser.Sid,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})
	resp := fiber.Map{
		"username":  newUser.Username,
		"sid":       newUser.Sid,
		"ipaddress": newUser.Ipaddress,
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user registered successfully", "user": resp})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	cookie := c.Cookies("my_cookie")
	if cookie == "" {
		return c.SendString("No cookie found")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "my_cookie",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		MaxAge:   -1,
		HTTPOnly: true,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "user logged out successfully"})
}
