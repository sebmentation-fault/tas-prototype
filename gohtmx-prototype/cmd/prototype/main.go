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
		File: "./pkg/assets/favicon.ico",
		URL:  "/favicon.ico",
	}))

	// In HTML code, we use '/static' to reach the assets in 'pkg/assets/'
	app.Static("/static", "./pkg/assets/")

	// Handle get requests to '/'
	app.Get("/", handlers.Index)

	// If no requests succeeds, then show the 404 not found page
	app.Use(handlers.NotFoundMiddleware)

	app.Listen(":42069")
}
