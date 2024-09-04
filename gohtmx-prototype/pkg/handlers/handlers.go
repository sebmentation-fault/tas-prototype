package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/supabase-go"
)

func SetupHandlers(app *fiber.App, client *supabase.Client) {
	// Handle get requests to '/event'
	app.Get("/event", eventHandler)

	// Handle get requests to '/dashboard'
	app.Get("/dashboard", dashboardHandler)
	app.Get("/dashboard/events", eventsHandler)
	app.Get("/dashboard/selfies", feedHandler)

	// Handle get requests to '/auth'
	app.Get("/auth/signup", signUpHandler)
	app.Get("/auth/login", logInHandler)

	// From the form submits
	app.Post("/auth/signup", signUpSubmittedHandlerWrapper(client))
	app.Post("/auth/login", logInSubmittedHandlerWrapper(client))

	// Handle get requests to '/'
	app.Get("/", Index)

	// If no requests succeeds, then show the 404 not found page
	app.Use(notFoundMiddleware)
}
