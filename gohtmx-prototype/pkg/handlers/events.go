package handlers

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	authservices "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/auth-services"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/db"
	eventss "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/components/events"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts/base"
)

func SetupEventsHandlers(server *TASServer) {
	app := server.Server
	events := app.Group("/events")
	events.Use(authservices.AuthMiddleware)

	// Getting the events can take a while.
	// We instantly return the skeleton events dashboard structure
	// Then we actually make a DB call and fill it up
	// the dashbaord (skeleton)
	events.Get("/", func(c *fiber.Ctx) error {
		templ := base.Base("Events Dashboard", "Events Dashboard", nil, eventss.EventsSkeleton())
		return renderHTML(c, templ)
	}) // /events
	events.Get("data", func(c *fiber.Ctx) error {
		// only the events is needed here
		// we do not need to send over the header/footer/etc again

		// for now lets just use this mock data
		sections := map[string][]db.Event{
			"Events Near by":      nearEvents,
			"Events Further away": nearEvents,
		}

		return renderHTML(c, eventss.Events(sections))
	}) // /events/data

	// an individual event
	events.Get("/:event_id<min(0)>", func(c *fiber.Ctx) error {
		eventId := c.Params("event_id")
		slog.Info("Request to event " + eventId)

		templ := base.Base("Event Infomation", "Event Information", nil, eventss.EventSkeleton("/events/data/1"))
		return renderHTML(c, templ)
	}) // /events/:event_id<min(0)>
	events.Get("/data/:event_id<min(0)>", func(c *fiber.Ctx) error {
		event := db.Event{
			ID:           1,
			CelebrityID:  "1",
			Title:        "Selfie",
			When:         time.Now(),
			Description:  "Very good description here.",
			IsReservedBy: nil,
			IsDeleted:    false,
			Price:        "",
			Location:     "Cafe",
			City:         "London",
			Country:      "United Kingdom",
		}

		return renderHTML(c, eventss.Event(&event))
	})
}

var (
	nearEvents = []db.Event{
		{
			ID:           1,
			CelebrityID:  "1",
			Title:        "Selfie",
			When:         time.Now(),
			Description:  "Very good description here.",
			IsReservedBy: nil,
			IsDeleted:    false,
			Price:        "",
			Location:     "Cafe",
			City:         "London",
			Country:      "United Kingdom",
		},
		{
			ID:           1,
			CelebrityID:  "1",
			Title:        "Selfie",
			When:         time.Now(),
			Description:  "Very good description here.",
			IsReservedBy: nil,
			IsDeleted:    false,
			Price:        "",
			Location:     "Cafe",
			City:         "London",
			Country:      "United Kingdom",
		}, {
			ID:           1,
			CelebrityID:  "1",
			Title:        "Selfie",
			When:         time.Now(),
			Description:  "Very good description here.",
			IsReservedBy: nil,
			IsDeleted:    false,
			Price:        "",
			Location:     "Cafe",
			City:         "London",
			Country:      "United Kingdom",
		}, {
			ID:           1,
			CelebrityID:  "1",
			Title:        "Selfie",
			When:         time.Now(),
			Description:  "Very good description here.",
			IsReservedBy: nil,
			IsDeleted:    false,
			Price:        "",
			Location:     "Cafe",
			City:         "London",
			Country:      "United Kingdom",
		}, {
			ID:           1,
			CelebrityID:  "1",
			Title:        "Selfie",
			When:         time.Now(),
			Description:  "Very good description here.",
			IsReservedBy: nil,
			IsDeleted:    false,
			Price:        "",
			Location:     "Cafe",
			City:         "London",
			Country:      "United Kingdom",
		},
	}
)
