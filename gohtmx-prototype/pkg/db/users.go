package db

type UserType int

const (
	AdminUser UserType = iota
	FanUser
	CelebrityUser
)

// A User could be admin, fan or celebrity
//
// CelebInfo only availiable if the user is a celebrity (duh)
type User interface {
	GetUserType() UserType
}

// Admin represents an admin user
type Admin struct {
	Account *Account
}

func (a *Admin) GetUserType() UserType {
	return AdminUser
}

// Fan represents a fan user
type Fan struct {
	Account *Account
}

func (f *Fan) GetUserType() UserType {
	return FanUser
}

// Celeb represents a celebrity user
type Celeb struct {
	Account   *Account
	CelebInfo *Celebrity
}

func (c *Celeb) GetUserType() UserType {
	return CelebrityUser
}
