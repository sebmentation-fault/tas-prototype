package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/handlers"
)

func main() {
	app := fiber.New()

	// Middleware to cache the favicon, speeding up acces to favicon requests
	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
		URL:  "/favicon.ico",
	}))

	// On client-leveL code, we use '/static' to reach the assets in 'public/',
	// to access CSS, JS, etc.
	app.Static("/static", "./public/")

	// Set up the handlers
	handlers.SetupHandlers(app)

	// Start listening
	app.Listen(":8080")
}
