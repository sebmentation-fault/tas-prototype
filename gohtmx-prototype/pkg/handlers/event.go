package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/events"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
)

// Get the Event details page
func eventHandler(c *fiber.Ctx) error {
	return RenderHTML(c, layouts.Base("Event Page", events.Event()))
}
