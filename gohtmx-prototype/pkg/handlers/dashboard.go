package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/accounts"
	services_events "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/events"
	view_events "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/events"
	view_layouts "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
)

// Get the Dashboard page
func dashboardHandler(c *fiber.Ctx) error {
	// TODO: get account
	var acc = &accounts.Account{
		Id:            "0",
		DisplayName:   "John Doe",
		IsDeleted:     false,
		IsCelebrity:   false,
		DateCreated:   "01/01/2001",
		DateLastLogin: "15/09/2024",
	}

	return RenderHTML(c, view_layouts.AuthorizedBase("Dashboard", acc, view_events.Tabs(), view_events.Events()))
}

// Get the events list skeleton
//
// Has no layout
func eventsHandler(c *fiber.Ctx) error {
	return RenderHTML(c, view_events.Events())
}

// Get the actual events sections
func eventSectionsHandler(c *fiber.Ctx) error {
	// just for testing the time pretending like network is slow
	time.Sleep(time.Second * 1)

	es, err := services_events.GetEventsBySection()
	if err != nil {
		// TODO change to server failture or smt (code 500)
		c.Set("HX-Redirect", "/404-not-found")
		return c.SendString("redirecting...")
	}

	return RenderHTML(c, view_events.EventsFilled(es))
}

// Get the global feed
//
// Has no layout
func feedHandler(c *fiber.Ctx) error {
	return RenderHTML(c, view_events.Feed())
}
