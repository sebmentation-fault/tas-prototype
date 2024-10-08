package db

import (
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/assert"
)

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
	GetID() int64
	GetUserType() UserType
	GetName() string
}

// Admin represents an admin user
type Admin struct {
	Account *Account
}

func (a *Admin) GetID() int64 {
	assert.NotNil(a, "There is no admin object")
	assert.NotNil(a.Account, "There is no admin account object")
	return a.Account.ID
}

func (a *Admin) GetUserType() UserType {
	return AdminUser
}

func (a *Admin) GetName() string {
	assert.NotNil(a, "There is no admin object")
	assert.NotNil(a.Account, "There is no admin account object")
	return a.Account.Username
}

// Fan represents a fan user
type Fan struct {
	Account *Account
}

func (f *Fan) GetID() int64 {
	assert.NotNil(f, "There is no fan object")
	assert.NotNil(f.Account, "There is no fan account object")
	return f.Account.ID
}

func (f *Fan) GetUserType() UserType {
	return FanUser
}

func (f *Fan) GetName() string {
	assert.NotNil(f, "There is no fan object")
	assert.NotNil(f.Account, "There is no fan account object")
	return f.Account.Username
}

// Celeb represents a celebrity user
type Celeb struct {
	Account   *Account
	CelebInfo *Celebrity
}

func (c *Celeb) GetID() int64 {
	assert.NotNil(c, "There is no celebrity object")
	assert.NotNil(c.Account, "There is no celebrity account object")
	return c.Account.ID
}

func (c *Celeb) GetUserType() UserType {
	return CelebrityUser
}

func (c *Celeb) GetName() string {
	assert.NotNil(c, "There is no celebrity object")
	assert.NotNil(c.Account, "There is no celebrity account object")
	return c.Account.Username
}
