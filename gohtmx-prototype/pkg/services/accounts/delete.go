package accounts

import (
	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

// Delete an account
//
// Effect: sets the auth.users.deleted_at field to the current time
func DeleteAccount(s *supabase.Client, a *Account) error {
	// Note: must parse can panic if string is not a uuid
	uuid := uuid.MustParse(a.Id)
	req := types.AdminDeleteUserRequest{
		UserID: uuid,
	}

	err := s.Auth.AdminDeleteUser(req)
	if err != nil {
		return err
	}

	a.IsDeleted = true
	return nil
}
