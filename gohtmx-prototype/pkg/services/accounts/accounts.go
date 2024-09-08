package accounts

const (
	CelebritiesTableName = "celebrities"
	ColumnCelebrityId    = "celebrity_id"
	ColumnCreatedAt      = "created_at"
)

// The struct for an account
type Account struct {
	Id          string `json:"user_id"`
	DisplayName string `json:"display_name"`

	IsDeleted   bool `json:"is_deleted"`
	IsCelebrity bool `json:"is_celebrity"`

	DateCreated   string `json:"date_created"`
	DateLastLogin string `json:"date_last_login"`
}

// Make a NewAccount instance
func NewAccount(id, name string, del, cel bool, create, login string) *Account {
	return &Account{
		Id:          id,
		DisplayName: name,

		IsDeleted:   del,
		IsCelebrity: cel,

		DateCreated:   create,
		DateLastLogin: login,
	}
}
