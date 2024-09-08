package account

const (
	ColumnCreatedAt = "created_at"
)

type AccountIdType string

// The account structure (fan/celeb, who knows)
type Account struct {
	Id          AccountIdType `json:"account_id"`
	DisplayName string        `json:"display_name"`
}

// Make a NewCelebrity instance
func NewAccount(id AccountIdType, n string) *Account {
	return &Account{
		Id:          id,
		DisplayName: n,
	}
}
