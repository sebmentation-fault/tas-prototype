package celebrities

import (
	"errors"

	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/assert"
	"github.com/supabase-community/supabase-go"
)

// Get celebrity information by id
func ReadCelebrity(s *supabase.Client, id CelebrityIdType) (*Celebrity, error) {
	assert.NotNil(s, "[GetCelebrity] supabase client is nil")
	assert.NotEmpty(string(id), "[GetCelebrity] celebrity id is empty")

	// why 'exact'? idk it was in the docs tho
	res, i, err := s.From(TableName).Select(ColumnCelebrityId, "exact", false).Execute()
	if err != nil {
		return nil, err

	}

	var _ = res
	var _ = i

	return nil, errors.New("todo")
}

// Get all the celebrities
func ReadCelebrities(s *supabase.Client) ([]Celebrity, error) {
	return nil, errors.New("todo")
}
