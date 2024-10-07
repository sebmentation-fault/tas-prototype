package db

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the Accounts table
type Account struct {
	ID                   int64     `db:"id"                      json:"id"`                      // User's Id
	Type                 int       `db:"type"                    json:"type"`                    // Enum to user type (consider changing to 'role')
	Username             string    `db:"username"                json:"username"`                // Unique username
	Email                string    `db:"email"                   json:"email"`                   // Unique email
	HashedPassword       string    `db:"hashed_password"         json:"hashed_password"`         // Hashed password
	DateTimeJoined       time.Time `db:"datetime_joined"         json:"datetime_joined"`         // Date and time joined
	DateTimeLastLoggedIn time.Time `db:"datetime_last_logged_in" json:"datetime_last_logged_in"` // Date and time last logged in
}

// Celebrity represents a celebrity in the Celebrities table
type Celebrity struct {
	AccountID int64  `db:"account_id"   json:"account_id"` // Foreign key to Accounts table
	Biography string `db:"biography" json:"biography"`     // Biography of the celebrity
}

// Event represents an event in the Events table
type Event struct {
	ID           int64     `db:"id"             json:"id"`             // Event Id
	CelebrityID  string    `db:"celebrity_id"   json:"celebrity_id"`   // Foreign key to Celebrities table
	Title        string    `db:"title"          json:"title"`          // Title
	When         time.Time `db:"happens_when"           json:"when"`   // Time that it happens
	Description  string    `db:"description"    json:"description"`    // Event description
	IsReservedBy *string   `db:"is_reserved_by" json:"is_reserved_by"` // Nullable reference to Accounts table - if null then not reserved
	IsDeleted    bool      `db:"is_deleted"     json:"is_deleted"`     // Indicates if the event is deleted
	Price        string    `db:"price"          json:"price"`          // Price
	Location     string    `db:"location"       json:"location"`       // Location
	City         string    `db:"city"           json:"city"`           // City
	Country      string    `db:"country"        json:"country"`        // Country
}

const (
	createSchema = `
CREATE TABLE IF NOT EXISTS Accounts (
	id                      INTEGER  PRIMARY KEY AUTOINCREMENT,
	type                    INT      NOT NULL,
	username                TEXT     NOT NULL    UNIQUE,
	email                   TEXT     NOT NULL    UNIQUE,
	hashed_password         TEXT     NOT NULL,
	datetime_joined         DATETIME NOT NULL    DEFAULT CURRENT_TIMESTAMP,
	datetime_last_logged_in DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS Celebrities (
	account_id INTEGER PRIMARY KEY REFERENCES Accounts(id) ON DELETE CASCADE,
	biography  TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS Events (
	id             INTEGER  PRIMARY KEY AUTOINCREMENT,
	celebrity_id   TEXT     NOT NULL    REFERENCES Celebrities(id) ON DELETE CASCADE,
	title          TEXT     NOT NULL,
	happens_when   DATETIME NOT NULL    DEFAULT CURRENT_TIMESTAMP,
	description    TEXT     NOT NULL,
	is_reserved_by INTEGER  REFERENCES Accounts(id) ON DELETE SET NULL,
	is_deleted     BOOLEAN  NOT NULL    DEFAULT FALSE,
	price          TEXT     NOT NULL,
	location       TEXT     NOT NULL,
	city           TEXT     NOT NULL,
	country        TEXT     NOT NULL
);
	`

	dropSchema = `
DROP TABLE Accounts;
DROP TABLE Celebrities;
DROP TABLE Events;
	`
)

func NewDatabase() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "takeaselfie.db")
	if err != nil {
		log.Panicln(err)
	}

	// create the schema
	_ = db.MustExec(createSchema)

	// create an admin user
	adminType := AdminUser
	adminUser := "admin"
	hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		panic("could not hash the admin password")
	}
	adminEmail := "admin@takeaselfie.com"
	lastLogin := time.Now()
	query := "INSERT OR IGNORE INTO Accounts (type, username, email, hashed_password, datetime_last_logged_in) VALUES (?, ?, ?, ?, ?)"
	_, err = db.
		Exec(query, adminType, adminUser, adminEmail, hash, lastLogin)
	if err != nil {
		panic("could not create admin")
	}

	return db
}
