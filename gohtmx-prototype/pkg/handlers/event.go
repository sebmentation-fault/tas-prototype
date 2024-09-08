package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	event_service "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/events"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/events"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
)

// Get the Event details page with skeletons
func eventHandler(c *fiber.Ctx) error {
	id := c.Params("event_id")
	url := "/event/info/" + id
	return RenderHTML(c, layouts.Base("Event Page", events.Event(url)))
}

// Get the Event details page filled
func eventFilledHandler(c *fiber.Ctx) error {
	// mock a slow network
	time.Sleep(time.Second * 2)

	id := c.Params("event_id")
	e, err := event_service.GetEvent(id)

	if err != nil {
		// TODO change to server failture or smt (code 500)
		c.Set("HX-Redirect", "/404-not-found")
		return c.SendString("redirecting...")
	}

	return RenderHTML(c, events.EventFilled(e, e.GetActivity().ImageURL))
}
