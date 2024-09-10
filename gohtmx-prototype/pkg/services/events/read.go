package events

import (
	"errors"

	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/accounts"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/services/activities"
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
		{
			Title:  "Upcoming, close",
			Events: mockUpcomingEvents,
		},
		{
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
func GetEvent(id string) (*Event, error) {

	// find the event in mockUpcomingEvents with matching id
	for _, e := range mockUpcomingEvents {
		if e.Id == id {
			return &e, nil
		}
	}

	return nil, errors.New("Could not find the event with the given ID")
}

var mockUpcomingEvents = []Event{
	*NewEvent(
		"0",
		false,
		false,
		"Cafe",
		"Have a coffee with me.",
		"£100",
		"London",
		"United Kingdom",
		accounts.NewAccount("0", "John Doe", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeCafe),
	),
	*NewEvent(
		"1",
		false,
		false,
		"Walk in Regents Park",
		"Meet with me where we go on a lovely walk for an hour or so.",
		"£10",
		"London",
		"United Kingdom",
		accounts.NewAccount("1", "A Famous Celeb", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeWalk),
	),
	*NewEvent(
		"2",
		false,
		false,
		"Cafe",
		"Have a coffee with me, again!.",
		"",
		"London",
		"United Kingdom",
		accounts.NewAccount("0", "John Doe", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeCafe),
	),
	*NewEvent(
		"3",
		false,
		false,
		"Book reading",
		"Have a book read it would be fun i think maybe.",
		"£20.00",
		"London",
		"United Kingdom",
		accounts.NewAccount("2", "A mum", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeChat),
	),
	*NewEvent(
		"4",
		false,
		false,
		"Hike",
		"Lets hike mount snowdonia in wales.",
		"£5.00",
		"mount snowdoniaiaia",
		"Wales",
		accounts.NewAccount("3", "Hiker-man 123", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeCafe),
	),
}

var mockUpcomingEventsFurtherAway = []Event{
	*NewEvent(
		"5",
		false,
		false,
		"Drinks at pub",
		"Have a drink with me at this cosy pub.",
		"£200",
		"Shottingham",
		"United Kingdom",
		accounts.NewAccount("4", "Drinking girl", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypePub),
	),
	*NewEvent(
		"6",
		false,
		false,
		"Title",
		"Very descriptive description of what is going to happen blah blah blah",
		"£10",
		"London",
		"United Kingdom",
		accounts.NewAccount("1", "A Famous Celeb", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeWalk),
	),
	*NewEvent(
		"7",
		false,
		false,
		"Title",
		"Very descriptive description of what is going to happend blah blah blah",
		"£1000",
		"Leeds",
		"United Kingdom",
		accounts.NewAccount("0", "John Doe", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeCafe),
	),
	*NewEvent(
		"8",
		false,
		false,
		"Title",
		"AAAAAAAAAAAAAAAAAAAAAAAAAA asdfasdfasdf adsf asdf sadf asdf asdf asdf asdf asdf asdf asdfasdf asdf asdf asdf asdf",
		"£20.00",
		"City",
		"Far Far Country",
		accounts.NewAccount("2", "A mum", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeChat),
	),
	*NewEvent(
		"9",
		false,
		false,
		"title",
		"qweryuiop asdf ghjkl;",
		"lots",
		"city",
		"country",
		accounts.NewAccount("3", "Hiker-man 123", false, true, "", ""),
		activities.NewActivity(activities.ActivityTypeHike),
	),
}
