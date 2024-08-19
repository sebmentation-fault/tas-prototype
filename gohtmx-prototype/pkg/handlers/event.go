package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/events"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
)

// Get the Event details page
func Event(c *fiber.Ctx) error {
	return RenderHTML(c, layouts.Base(events.Event()))
}
