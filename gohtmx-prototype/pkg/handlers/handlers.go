package handlers

import (
	"fmt"

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
	// TODO: change this to an actual secret (e.g. in an ENV)
	signingKey := []byte("secret")
	app.Post("/auth/signup", signUpSubmittedHandlerWrapper(client, signingKey))
	app.Post("/auth/login", logInSubmittedHandlerWrapper(client, signingKey))

	// ----- Authorized routes -----
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: signingKey},
		// read the cookie
		TokenLookup: "cookie:token",
		// if success, go the thing
		SuccessHandler: func(c *fiber.Ctx) error {
			fmt.Println("Successing")
			return c.Next() // Proceed if token is valid
		},
		// if something wrong, then we should address it
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// TODO: figure out actual error
			// * JWT token expired -> re-login
			// * straight up not present -> login
			// * is there a different error where something else should happen?
			// maybe requires adding message/toast/notification explaining message
			fmt.Println("Erroring: ", err)
			return c.Redirect("/auth")
		},
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
