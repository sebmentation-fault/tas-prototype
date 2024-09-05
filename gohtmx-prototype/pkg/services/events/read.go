package events

import "errors"

// TODO:
// Service to get some of the recent events from the supabase server
func GetEvents() ([]Event, error) {
	var events = []Event{}
	return events, errors.New("todo")
}

// TODO:
// Service to get an event by ID
func GetEvent(id int) (*Event, error) {
	return nil, errors.New("todo")
}
