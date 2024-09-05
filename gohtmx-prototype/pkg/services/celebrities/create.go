package celebrities

import (
	"errors"

	"github.com/supabase-community/supabase-go"
)

// A new celebrity signs up!! Very good news when this happens :)
//
// The celebrity is a normal supabase user, but their ids are in
// the celebrity table.
//
// # All that happens is we add the user id to the celebrities table
//
// If they could not be added, returns an error
func CreateCelebrity(s *supabase.Client, c *Celebrity) error {
	return errors.New("TODO")
}
