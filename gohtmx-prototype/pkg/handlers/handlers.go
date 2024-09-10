package handlers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/supabase-go"
)

func SetupHandlers(app *fiber.App, client *supabase.Client) {
	// ----- Unauthorized routes -----

	// Handle get requests to '/'
	app.Get("/", Index)

	// Handle get requests to '/auth'
	app.Get("/auth", authHandler)
	app.Get("/auth/signup", signUpHandler)
	app.Get("/auth/login", logInHandler)

	// From the form submits
	app.Post("/auth/signup", signUpSubmittedHandlerWrapper(client))
	app.Post("/auth/login", logInSubmittedHandlerWrapper(client))

	// ----- Authorized routes -----
	// TODO: change this to an actual secret
	signingKey := []byte("secret")
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: signingKey},
	}))

	// Handle get requests to '/event'
	app.Get("/event/:event_id", eventHandler)
	app.Get("/event/info/:event_id", eventFilledHandler)

	// Handle get requests to '/dashboard'
	app.Get("/dashboard", dashboardHandler)
	app.Get("/dashboard/events", eventsHandler)
	app.Get("/dashboard/event-sections", eventSectionsHandler)
	app.Get("/dashboard/selfies", feedHandler)
	// Handle get requests to '/headers'
	app.Get("/headers/account-info", accountButtonHandler)

	// If no requests succeeds, then show the 404 not found page
	app.Use(notFoundMiddleware)
}
