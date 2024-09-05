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

// TODO: add time created field (is already in the database column)

// The event struct
//
// If the event is being created, then there will be no Id, as this is assigned
// by the database.
// Therefore, there is only event ids on events that have been created/fetched
type Event struct {
	id        EventIdType
	isDeleted bool

	celebrity  *celebrities.Celebrity
	isReserved bool
	price      PriceType

	title       string
	description string

	activity *activities.Activity

	city    string
	country string
}

// Create a new event with the id, the celebrity, the resevation status, the
// price, the title, the description, the activity, the city and the country
//
// This function is used when creating the event from the database call.
// If creating an event from a form, then use the builder
func NewEvent(
	id EventIdType,
	dl bool,
	c *celebrities.Celebrity,
	r bool,
	p PriceType,
	t string,
	de string,
	a *activities.Activity,
	ci string,
	co string,
) *Event {
	return &Event{
		id:          id,
		isDeleted:   dl,
		celebrity:   c,
		isReserved:  r,
		price:       p,
		title:       t,
		description: de,
		activity:    a,
		city:        ci,
		country:     co,
	}
}

// Get the Id of the event
func (e *Event) GetId() EventIdType {
	return e.id
}

// Get the status for whether the event has been deleted
func (e *Event) GetDeletionStatus() bool {
	return e.isDeleted
}

// Get the celebrity of the event
func (e *Event) GetCelebrity() *celebrities.Celebrity {
	return e.celebrity
}

// Get the resevation status of the event
func (e *Event) GetResevationStatus() bool {
	return e.isReserved
}

// Get the price of the event
func (e *Event) GetPrice() PriceType {
	return e.price
}

// Get the title of the event
func (e *Event) GetTitle() string {
	return e.title
}

// Get the description of the event
func (e *Event) GetDescription() string {
	return e.description
}

// Get the activity of the event
func (e *Event) GetActivity() *activities.Activity {
	return e.activity
}

// Get the city of the event
func (e *Event) GetCity() string {
	return e.city
}

// Get the country of the event
func (e *Event) GetCountry() string {
	return e.country
}

// EventBuilder is a struct for building an event with the builder pattern
// because enterprise level shit
type EventBuilder struct {
	event *Event
}

// NewEventBuilder creates a new EventBuilder instance
//
// The required values are given defaults here
func NewEventBuilder() *EventBuilder {
	return &EventBuilder{
		event: &Event{},
	}
}

// WithCelebrity sets the celebrity of the event
func (eb *EventBuilder) WithCelebrity(c *celebrities.Celebrity) *EventBuilder {
	eb.event.celebrity = c
	return eb
}

// WithReserved sets the reservation status of the event
func (eb *EventBuilder) WithReserved(r bool) *EventBuilder {
	eb.event.isReserved = r
	return eb
}

// WithPrice sets the price of the event
func (eb *EventBuilder) WithPrice(p PriceType) *EventBuilder {
	eb.event.price = p
	return eb
}

// WithTitle sets the title of the event
func (eb *EventBuilder) WithTitle(t string) *EventBuilder {
	eb.event.title = t
	return eb
}

// WithDescription sets the description of the event
func (eb *EventBuilder) WithDescription(d string) *EventBuilder {
	eb.event.description = d
	return eb
}

// WithActivityType sets the activity of the event
func (eb *EventBuilder) WithActivityType(a activities.ActivityType) *EventBuilder {
	eb.event.activity = activities.NewActivity(a)
	return eb
}

// WithCity sets the city of the event
func (eb *EventBuilder) WithCity(c string) *EventBuilder {
	eb.event.city = c
	return eb
}

// WithCountry sets the country of the event
func (eb *EventBuilder) WithCountry(c string) *EventBuilder {
	eb.event.country = c
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
	appendErrIfTrue(eb.event.celebrity == nil, errs, NoCelebId)
	appendErrIfTrue(eb.event.title == "", errs, NoTitle)
	appendErrIfTrue(eb.event.description == "", errs, NoDescription)
	appendErrIfTrue(eb.event.city == "", errs, NoCity)
	appendErrIfTrue(eb.event.country == "", errs, NoCountry)

	// Return an error for each and every required value missed
	if len(errs) > 0 {
		return nil, errs
	}

	// Now we fill in the defaults :)
	if eb.event.price == "" {
		eb.event.price = "just the selfie :)"
	}

	if eb.event.activity == nil {
		eb.event.activity = activities.NewActivity(activities.ActivityTypeDefault)
	}

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
