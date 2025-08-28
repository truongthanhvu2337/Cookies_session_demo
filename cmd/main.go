package main

import (
	"cookies_session_demo/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	authHandler := new(services.AuthHandler)
	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)
	app.Post("/logout", authHandler.Logout)

	app.Listen(":3000")
}
