package accounts

import "github.com/supabase-community/supabase-go"

// Grant an account celebrity status
//
// Effect: inserts the account into the celebrities table
func GrantAccountAsCelebrity(s *supabase.Client, c *Account) error {
	_, _, err := s.From(CelebritiesTableName).Insert(c, false, "", "*", "").Execute()
	if err != nil {
		return err
	}

	c.IsCelebrity = true
	return nil
}

// Revoke the celebrity status of an account
//
// Effect: deletes the account from the celebrities table
func RevokeAccountFromCelebrity(s *supabase.Client, a *Account) error {
	// Delete the account from the celebrities table
	// where the celebrity_id is the accounts id
	_, _, err := s.From(CelebritiesTableName).Delete("", "").Eq(ColumnCelebrityId, a.Id).Execute()

	if err != nil {
		return err
	}

	a.IsCelebrity = false
	return nil
}
