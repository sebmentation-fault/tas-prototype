package handlers

import (
	"log/slog"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	authservices "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/auth-services"
	ddbb "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/db"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/components/auth/login"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/components/auth/signup"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts/base"
)

func SetupAuthHandlers(server *TASServer) {
	app := server.Server
	auth := app.Group("/auth")
	db := server.DB

	// post the forms
	auth.Post("/login", postLoginForm(db))   // /auth/login
	auth.Post("/signup", postSignupForm(db)) // /auth/signup

	// get the layout that the forms sit on
	auth.Get("/", func(c *fiber.Ctx) error {
		// by default, shows the signup one, user manually goes from there to login if need
		// Why?
		// I do not want to refresh the whole page when they do signup -> login
		// Instead, let's just change the form element using fancy HTMX
		return renderHTML(c, base.Base("Log in/Sign up", "Log in and Sign up", login.LogInTempl()))
	}) // /auth

	// get the components for the forms
	auth.Get("/login", func(c *fiber.Ctx) error {
		return renderHTML(c, login.LogInTempl())
	}) // /auth/login
	auth.Get("/signup", func(c *fiber.Ctx) error {
		return renderHTML(c, signup.SignUpTempl())
	}) // /auth/signup
}

// if unsuccessful, return an error message that updates the form's message
// box.
//
// if successful:
//   - generate a jwt claim
//   - store in the users cookie (yum yum)
//   - redirect to dashboard
func postLoginForm(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slog.Info("Login attempt")

		// testing to see the loading indicator :)
		time.Sleep(1 * time.Second)

		// get the email -- remove trailing (and leading) whitespace
		email := strings.TrimSpace(c.FormValue("email"))
		// get the password
		password := c.FormValue("password")

		// check if user exists
		query := "SELECT id, type, username, email, hashed_password, datetime_joined, datetime_last_logged_in FROM Accounts WHERE email = ?"
		rows, err := db.Queryx(query, email)
		if err != nil {
			slog.Error("Logging in user ", email, " returns error ", err)
			return renderHTML(c, login.ErrOnLogIn(fiber.ErrInternalServerError))
		}
		defer rows.Close()

		for rows.Next() {
			var account ddbb.Account
			err := rows.StructScan(&account)
			if err != nil {
				slog.Error("Could not scan into account struct: ", err)
				return renderHTML(c, login.ErrOnLogIn(fiber.ErrInternalServerError))
			}
			// Check email, then password
			if strings.Compare(account.Email, email) != 0 {
				return renderHTML(c, login.ErrOnLogIn(fiber.ErrUnauthorized))
			}
			if !authservices.ComparePasswordWithHash(password, account.HashedPassword) {
				return renderHTML(c, login.ErrOnLogIn(fiber.ErrUnauthorized))
			}

			// generate the jwt claim.
			signedToken, err := authservices.NewSignedJWTTokenWithClaims(&account)

			if err != nil {
				slog.Warn("Could not sign a token")
				return renderHTML(c, login.ErrOnLogIn(fiber.ErrInternalServerError))
			}

			// make a cookie (yum yum)
			cookie := &fiber.Cookie{
				Name:     "jwt",
				Value:    signedToken,
				Expires:  time.Now().Add(72 * time.Hour),
				HTTPOnly: true,
				Secure:   false, // set to True when on the interwebs and not dev mode
			}
			c.Cookie(cookie)

			slog.Info("Login successful!")
			c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
			c.Set("HX-Redirect", "/dashboard")
			return c.SendString("Redirecting to /dashboard...")
		}

		slog.Info("Login attempt failed - no rows to iterate over")
		return renderHTML(c, login.ErrOnLogIn(fiber.ErrUnauthorized))
	}
}

// if unsuccessful, return an error message that updates the form's message
// box.
//
// if successful:
//   - generate a jwt claim
//   - store in the users cookie (yum yum)
//   - redirect to dashboard
func postSignupForm(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slog.Info("Signup attempt")

		// testing to see the loading indicator :)
		time.Sleep(1 * time.Second)

		// get the username -- remove trailing (and leading) whitespace
		username := strings.TrimSpace(c.FormValue("username"))
		// get the email -- remove trailing (and leading) whitespace
		email := strings.TrimSpace(c.FormValue("email"))
		// get the password
		password := c.FormValue("password")

		// check if user exists
		rows, err := db.Queryx("SELECT * FROM Accounts WHERE email = ? OR username = ?", email, username)
		if err != nil {
			slog.Error("Sign up failed because ", err)
			return renderHTML(c, signup.ErrOnSignUp(fiber.ErrInternalServerError))
		}
		defer rows.Close()

		for rows.Next() {
			slog.Info("Signing up does not work if the user email/username already exists")
			return renderHTML(c, signup.ErrOnSignUp(fiber.ErrConflict))
		}

		// user does not exist! lets do some signing up
		hash, err := authservices.GenerateHashFromPassword(password)
		if err != nil {
			slog.Error("Signup failed to generate a hash from the password: ", err)
			return renderHTML(c, signup.ErrOnSignUp(fiber.ErrInternalServerError))
		}

		// account needs:
		// we provide: type, username, email, hash, datelastlogged in,
		// ddbb provides: id, date created
		account := ddbb.Account{
			Type:                 int(ddbb.FanUser),
			Username:             username,
			Email:                email,
			HashedPassword:       hash,
			DateTimeLastLoggedIn: time.Now(),
		}
		query := "INSERT INTO Accounts (type, username, email, hashed_password, datetime_last_logged_in) VALUES (?, ?, ?, ?, ?) RETURNING id, datetime_joined"
		err = db.
			QueryRowx(query, account.Type, account.Username, account.Email, account.HashedPassword, account.DateTimeLastLoggedIn).
			Scan(&account.ID, &account.DateTimeJoined)
		if err != nil {
			slog.Warn("Could not insert fan account to db: ", err)
			return renderHTML(c, signup.ErrOnSignUp(fiber.ErrInternalServerError))
		}

		// generate the jwt claim.
		signedToken, err := authservices.NewSignedJWTTokenWithClaims(&account)

		if err != nil {
			slog.Warn("Could not sign a token")
			return renderHTML(c, signup.ErrOnSignUp(fiber.ErrInternalServerError))
		}

		// make a cookie (yum yum)
		cookie := &fiber.Cookie{
			Name:     "jwt",
			Value:    signedToken,
			Expires:  time.Now().Add(72 * time.Hour),
			HTTPOnly: true,
			Secure:   false, // set to True when on the interwebs and not dev mode
		}
		c.Cookie(cookie)

		slog.Info("Signup successful!")
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		c.Set("HX-Redirect", "/dashboard")
		return c.SendString("Redirecting to /dashboard...")
	}
}
