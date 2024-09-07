package events

import (
	"errors"

	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/assert"
	"github.com/supabase-community/supabase-go"
)

// TODO:
// Service to create an event
//
// Panics if the inputs are nil, because I'm cruel and require valid inputs.
// If inputs are nil and therefore invalid, I'm going to assume something is
// incredibly wrong with the code
func CreateEvent(s *supabase.Client, e *Event) error {
	assert.NotNil(s, "[CreateEvent] s is nil")
	assert.NotNil(e, "[CreateEvent] e is nil")

	// s.From()

	return errors.New("todo")
}
