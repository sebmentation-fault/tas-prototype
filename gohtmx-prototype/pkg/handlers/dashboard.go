package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/events"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
)

// Get the Dashboard page
func Dashboard(c *fiber.Ctx) error {
	return RenderHTML(c, layouts.Base(events.Tabs(), events.Events()))
}

// Get the events list
//
// Has no layout
func Events(c *fiber.Ctx) error {
	return RenderHTML(c, events.Events())
}

// Get the global feed
//
// Has no layout
func Feed(c *fiber.Ctx) error {
	return RenderHTML(c, events.Feed())
}
