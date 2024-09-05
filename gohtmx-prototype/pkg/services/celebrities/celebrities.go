package celebrities

const (
	TableName = "celebrities"
)

type CelebrityIdType string

// The struct for the celebrity
type Celebrity struct {
	id          CelebrityIdType
	displayName string
}

// Make a NewCelebrity instance
func NewCelebrity(id CelebrityIdType, n string) *Celebrity {
	return &Celebrity{
		id:          id,
		displayName: n,
	}
}

// Get the id for the celebrity
func (c *Celebrity) GetId() CelebrityIdType {
	return c.id
}

// Get the display name for the celebrity
func (c *Celebrity) GetDisplayName() string {
	return c.displayName
}
