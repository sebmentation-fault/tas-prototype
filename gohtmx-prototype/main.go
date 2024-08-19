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

	// Handle get requests to '/event'
	app.Get("/event", handlers.Event)
	// Handle get requests to '/dashboard'
	app.Get("/dashboard", handlers.Dashboard)
	app.Get("/dashboard/events", handlers.Events)
	app.Get("/dashboard/selfies", handlers.Feed)

	// Handle get requests to '/'
	app.Get("/", handlers.Index)

	// If no requests succeeds, then show the 404 not found page
	app.Use(handlers.NotFoundMiddleware)

	app.Listen(":42069")
}
