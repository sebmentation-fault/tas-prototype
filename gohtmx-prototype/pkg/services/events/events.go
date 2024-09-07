package events

import (
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/activities"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/celebrities"
)

const (
	TableName = "events"
)

type EventIdType string
type PriceType string

// TODO: add time created field (is already in the database column)?

// The event struct
//
// If the event is being created, then there will be no Id, as this is assigned
// by the database.
// Therefore, there is only event ids on events that have been created/fetched
type Event struct {
	Id          EventIdType            `json:"event_id"`
	Celebrity   *celebrities.Celebrity `json:"celebrity"`
	IsDeleted   bool                   `json:"is_deleted"`
	IsReserved  bool                   `json:"is_reserved"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Activity    *activities.Activity   `json:"activity"`
	Price       PriceType              `json:"price"`
	City        string                 `json:"city"`
	Country     string                 `json:"country"`
}

// Create a new event with the id, the celebrity, the resevation status, the
// price, the title, the description, the activity, the city and the country
//
// This function is used when creating the event from the database call.
// If creating an event from a form, then use the builder
func NewEvent(
	id EventIdType,
	celeb *celebrities.Celebrity,
	del bool,
	res bool,
	title string,
	des string,
	act *activities.Activity,
	p PriceType,
	city string,
	country string,
) *Event {
	return &Event{
		Id:          id,
		IsDeleted:   del,
		Celebrity:   celeb,
		IsReserved:  res,
		Title:       title,
		Description: des,
		Price:       p,
		Activity:    act,
		City:        city,
		Country:     country,
	}
}

// The struct for different sections of events
type EventSections struct {
	Title  string  `json:"title"`
	Events []Event `json:"events"`
}

// EventBuilder is a struct for building an event with the builder pattern
// because enterprise level shit is cool sometimes
type EventBuilder struct {
	event *Event
}

// NewEventBuilder creates a new EventBuilder instance
//
// The optional values are given defaults here
func NewEventBuilder() *EventBuilder {
	return &EventBuilder{
		event: &Event{
			Price:    "just the selfie :)",
			Activity: activities.NewActivity(activities.ActivityTypeDefault),
		},
	}
}

// WithCelebrity sets the celebrity of the event
func (eb *EventBuilder) WithCelebrity(c *celebrities.Celebrity) *EventBuilder {
	eb.event.Celebrity = c
	return eb
}

// WithReserved sets the reservation status of the event
func (eb *EventBuilder) WithReserved(r bool) *EventBuilder {
	eb.event.IsReserved = r
	return eb
}

// WithPrice sets the price of the event
func (eb *EventBuilder) WithPrice(p PriceType) *EventBuilder {
	eb.event.Price = p
	return eb
}

// WithTitle sets the title of the event
func (eb *EventBuilder) WithTitle(t string) *EventBuilder {
	eb.event.Title = t
	return eb
}

// WithDescription sets the description of the event
func (eb *EventBuilder) WithDescription(d string) *EventBuilder {
	eb.event.Description = d
	return eb
}

// WithActivityType sets the activity of the event
func (eb *EventBuilder) WithActivityType(a activities.ActivityType) *EventBuilder {
	eb.event.Activity = activities.NewActivity(a)
	return eb
}

// WithCity sets the city of the event
func (eb *EventBuilder) WithCity(c string) *EventBuilder {
	eb.event.City = c
	return eb
}

// WithCountry sets the country of the event
func (eb *EventBuilder) WithCountry(c string) *EventBuilder {
	eb.event.Country = c
	return eb
}

// Build returns the built Event, or a list of errors if the event has invalid
// fields.
//
// Use Build when using the form to create a new event and see if it can be created.
// Use NewEvent when creating a new event instance from having fetched it from the
// database.
//
// Required fields include:
// * the celebrity
// * the title
// * the description
// * the city and country
//
// Optional fields have the following defaults:
// * price - nothing! free!
// * isReserved - false!
// * the activity type - default activity
//
// Id is not availible in this method, because the database assigns the Id.
// IsDeleted is also not availiable, because the user should not be able to
// pre-create a deleted event
func (eb *EventBuilder) Build() (*Event, []error) {
	var errs []error

	// Check the required values
	appendErrIfTrue(eb.event.Celebrity == nil, errs, NoCelebId)
	appendErrIfTrue(eb.event.Title == "", errs, NoTitle)
	appendErrIfTrue(eb.event.Description == "", errs, NoDescription)
	appendErrIfTrue(eb.event.City == "", errs, NoCity)
	appendErrIfTrue(eb.event.Country == "", errs, NoCountry)

	// Return an error for each and every required value missed
	if len(errs) > 0 {
		return nil, errs
	}

	// Otherwise we return the build event :D
	return eb.event, nil
}

// If the condition is true, then append the error to the error list
//
// Helper function to clean up Build method
func appendErrIfTrue(cond bool, errs []error, err error) {
	if cond {
		errs = append(errs, err)
	}
}
