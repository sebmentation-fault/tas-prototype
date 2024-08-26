package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views"
	errorhandling "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/error-handling"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
)

// Index handler returns the Index page.
//
// If the user is authenticated, then the Index page is the user's dashboard.
// Otherwise, shows the TakeASelfie home page.
func Index(c *fiber.Ctx) error {
	return RenderHTML(c, layouts.Base(view.Index()))
}

// Middleware to show a 404 page-not-found page if the page was not found.
//
// From the templ docs.
func notFoundMiddleware(c *fiber.Ctx) error {
	// Set the HTTP status code
	c.Status(fiber.StatusNotFound)
	return RenderHTML(c, layouts.Error(errorhandling.NotFound()))
}

// A Render function that renders templ-HTML with a given Fiber context.
//
// From the templ docs.
func RenderHTML(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
