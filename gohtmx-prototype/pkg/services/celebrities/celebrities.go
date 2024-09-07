package celebrities

const (
	TableName         = "celebrities"
	ColumnCelebrityId = "celebrity_id"
	ColumnCreatedAt   = "created_at"
)

type CelebrityIdType string

// The struct for the celebrity
type Celebrity struct {
	Id          CelebrityIdType `json:"celebrity_id"`
	DisplayName string          `json:"display_name"`
}

// Make a NewCelebrity instance
func NewCelebrity(id CelebrityIdType, n string) *Celebrity {
	return &Celebrity{
		Id:          id,
		DisplayName: n,
	}
}
