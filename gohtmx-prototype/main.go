package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/db"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/handlers"
)

func main() {
	// Connect to la database
	db := db.NewDatabase()
	defer db.Close()

	// Set up the fiber app
	app := fiber.New()

	server := &handlers.TASServer{
		Server: app,
		DB:     db,
		Store:  nil,
	}

	// Middleware to cache the favicon, speeding up acces to favicon requests
	// Also corrects the URL
	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
		URL:  "/favicon.ico",
	}))

	// On client-leveL code, we use '/static' to reach the assets in 'public/',
	// to access CSS, JS, etc.
	app.Static("/static", "./public/")

	// Set up all the handlers for the different routes
	handlers.SetupHandlers(server)

	// Start listening
	app.Listen(":8080")
}
