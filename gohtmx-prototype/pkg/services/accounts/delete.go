package accounts

import (
	"errors"

	"github.com/supabase-community/supabase-go"
)

// Delete an account
//
// Effect: sets the auth.users.deleted_at field to the current time
func DeleteAccount(s *supabase.Client, a *Account) error {
	return errors.New("not implemented")
}
