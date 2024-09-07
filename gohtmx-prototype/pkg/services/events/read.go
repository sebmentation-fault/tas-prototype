package events

import (
	"errors"

	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/activities"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/celebrities"
)

// Service to get events by their different sections
// Also TODO maybe Cache these events
// where the cache expires/gets auto-updated whenever a celeb creates/changes an
// event.
// Then even later a different TODO would be only update the caches for local
// regions!! (e.g. someone in london should likely not need to be notified
// if there is a change to events in los angleleleles)
func GetEventsBySection() ([]EventSections, error) {
	var es = []EventSections{
		EventSections{
			Title:  "Upcoming, close",
			Events: mockUpcomingEvents,
		},
		EventSections{
			Title:  "Upcoming, further away",
			Events: mockUpcomingEventsFurtherAway,
		},
	}

	return es, nil
}

// TODO: stop using fake (ew) data and call the actual database
// Service to get some of the recent events from the supabase server
func GetEvents() ([]Event, error) {
	var events = mockUpcomingEvents

	return events, nil
}

// TODO:
// Service to get an event by ID
func GetEvent(id int) (*Event, error) {
	return nil, errors.New("todo")
}

var mockUpcomingEvents = []Event{
	*NewEvent(
		EventIdType("0"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("0"), "John Doe"),
		false,
		false,
		"Cafe",
		"Have a coffee with me.",
		activities.NewActivity(activities.ActivityTypeCafe),
		"",
		"London",
		"United Kingdom",
	),
	*NewEvent(
		EventIdType("1"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("1"), "A Famous Celeb"),
		false,
		false,
		"Walk in Regents Park",
		"Meet with me where we go on a lovely walk for an hour or so.",
		activities.NewActivity(activities.ActivityTypeWalk),
		"£10",
		"London",
		"United Kingdom",
	),
	*NewEvent(
		EventIdType("2"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("0"), "John Doe"),
		false,
		false,
		"Cafe",
		"Have a coffee with me, again!.",
		activities.NewActivity(activities.ActivityTypeCafe),
		"",
		"London",
		"United Kingdom",
	),
	*NewEvent(
		EventIdType("3"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("2"), "A mum"),
		false,
		false,
		"Book reading",
		"Have a book read it would be fun i think maybe.",
		activities.NewActivity(activities.ActivityTypeChat),
		"£20.00",
		"London",
		"United Kingdom",
	),
	*NewEvent(
		EventIdType("4"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("3"), "Hiker-man 123"),
		false,
		false,
		"Hike",
		"Lets hike mount snowdonia in wales.",
		activities.NewActivity(activities.ActivityTypeCafe),
		"£5.00",
		"mount snowdoniaiaia",
		"Wales",
	),
}

var mockUpcomingEventsFurtherAway = []Event{
	*NewEvent(
		EventIdType("5"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("4"), "Drinking girl"),
		false,
		false,
		"Drinks at pub",
		"Have a drink with me at this cosy pub.",
		activities.NewActivity(activities.ActivityTypePub),
		"£200",
		"Shottingham",
		"United Kingdom",
	),
	*NewEvent(
		EventIdType("6"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("1"), "A Famous Celeb"),
		false,
		false,
		"Title",
		"Very descriptive description of what is going to happen blah blah blah",
		activities.NewActivity(activities.ActivityTypeWalk),
		"£10",
		"London",
		"United Kingdom",
	),
	*NewEvent(
		EventIdType("7"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("0"), "John Doe"),
		false,
		false,
		"Title",
		"Very descriptive description of what is going to happend blah blah blah",
		activities.NewActivity(activities.ActivityTypeCafe),
		"£1000",
		"Leeds",
		"United Kingdom",
	),
	*NewEvent(
		EventIdType("8"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("2"), "A mum"),
		false,
		false,
		"Title",
		"AAAAAAAAAAAAAAAAAAAAAAAAAA asdfasdfasdf adsf asdf sadf asdf asdf asdf asdf asdf asdf asdfasdf asdf asdf asdf asdf",
		activities.NewActivity(activities.ActivityTypeChat),
		"£20.00",
		"City",
		"Far Far Country",
	),
	*NewEvent(
		EventIdType("9"),
		celebrities.NewCelebrity(celebrities.CelebrityIdType("3"), "Hiker-man 123"),
		false,
		false,
		"title",
		"qweryuiop asdf ghjkl;",
		activities.NewActivity(activities.ActivityTypeHike),
		"lots",
		"city",
		"country",
	),
}
