package handlers

import "github.com/gofiber/fiber/v2"

func SetupHandlers(app *fiber.App) {
	// Handle get requests to '/event'
	app.Get("/event", eventHandler)
	// Handle get requests to '/dashboard'
	app.Get("/dashboard", dashboardHandler)
	app.Get("/dashboard/events", eventsHandler)
	app.Get("/dashboard/selfies", feedHandler)

	// Handle get requests to '/'
	app.Get("/", Index)

	// If no requests succeeds, then show the 404 not found page
	app.Use(notFoundMiddleware)

}
