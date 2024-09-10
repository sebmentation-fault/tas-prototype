package events

import (
	"log/slog"

	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/assert"
	"github.com/supabase-community/supabase-go"
)

// Service to create an event
//
// Panics if the inputs are nil, because I'm cruel and require valid inputs.
// If inputs are nil and therefore invalid, I'm going to assume something is
// incredibly wrong with the code
func CreateEvent(s *supabase.Client, e *Event) error {
	assert.NotNil(s, "[CreateEvent] s is nil")
	assert.NotNil(e, "[CreateEvent] e is nil")
	assert.NotEmpty(e.Id, "[CreateEvent] event exists already (probably definitely)")

	_, _, err := s.From(TableName).Insert(e, false, "", "*", "").Execute()

	if err != nil {
		return err
	}

	if e.Id == "" {
		slog.Error("[CreateEvent] id is still empty after being made")
	} else {
		slog.Info("[CreateEvent] id updated successfully to %s", string(e.Id))
	}

	return nil
}
