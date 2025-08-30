package main

import (
	"cookies_session_demo/internal/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ProxyHeader: "X-Forwarded-For",
	})

	authHandler := new(services.AuthHandler)

	// If you want to trust proxy headers, set the ProxyHeader in the config (optional):
	// app.Config().ProxyHeader = "X-Forwarded-For"
	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)
	app.Post("/logout", authHandler.Logout)

	app.Listen(":3000")
}
