package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jmoiron/sqlx"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/components/hero"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts/base"
)

type TASServer struct {
	Server *fiber.App
	Store  *session.Store
	DB     *sqlx.DB
}

func SetupHandlers(server *TASServer) {
	app := server.Server

	// Handle get requests to '/'
	app.Get("/", func(c *fiber.Ctx) error {
		// check if user logged in, then the account should be passed in
		// (e.g. so that the header can be epic and show funky information)
		return renderHTML(c, base.Base("Get Started", "Index", nil, hero.Hero(nil)))
	})

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.SendString("The dashboard")
	})

	// set up the authentication-related handlers
	SetupAuthHandlers(server)
	SetupAdminHandlers(server)
	SetupEventsHandlers(server)
	SetupHeaderHandlers(server)
	// TODO: setup other handlers too

	// a not-found middleware
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})
}

// A Render function that renders templ-HTML with a given Fiber context.
//
// From the templ docs.
func renderHTML(c *fiber.Ctx, component templ.Component) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return component.Render(c.Context(), c.Response().BodyWriter())
}
