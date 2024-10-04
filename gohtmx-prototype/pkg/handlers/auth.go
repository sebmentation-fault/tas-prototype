package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/auth"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

// Default auth handler -- navigate to this one please
func authHandler(c *fiber.Ctx) error {
	return RenderHTML(c, layouts.AuthorizedBase("Authenticate", nil, auth.SignUpTempl()))
}

// the first two handlers simply render the html with the base layout

func signUpHandler(c *fiber.Ctx) error {
	return RenderHTML(c, auth.SignUpTempl())
}

func logInHandler(c *fiber.Ctx) error {
	return RenderHTML(c, auth.LogInTempl())
}

// these two handlers are going to respond to the log in/sign up form submition

func signUpSubmittedHandlerWrapper(s *supabase.Client, k []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return signUpSubmittedHandler(c, s, k)
	}
}

func signUpSubmittedHandler(c *fiber.Ctx, s *supabase.Client, k []byte) error {
	fmt.Println("trying to sign the up")

	// Parse the request body as JSON
	var signupData types.SignupRequest
	if err := c.BodyParser(&signupData); err != nil {
		return RenderHTML(c, auth.ErrOnSignUp(err))
	}

	res, err := s.Auth.Signup(signupData)

	if err != nil {
		// return RenderHTML(c, auth.ErrOnSignUp(err))
		// TODO: be descriptive with which error:
		// * user already exists
		// * supabase cannot be found?
		// * etc
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// user is signed up, we should automagically log them in to now
	claims := jwt.MapClaims{
		"name":  res.User,
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(k)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72), // Set expiration time
		HTTPOnly: true,                           // Prevents JavaScript access to the cookie -- in any case that we have scripting attacks, the attacker can not steal this information
		Secure:   false,                          // FIXME: Set to true if using HTTPS
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	c.Set("HX-Redirect", "/dashboard")
	return c.SendString("redirecting to /dashboard")
}

func logInSubmittedHandlerWrapper(s *supabase.Client, k []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return logInSubmittedHandler(c, s, k)
	}
}

func logInSubmittedHandler(c *fiber.Ctx, s *supabase.Client, k []byte) error {
	// Parse the request body as JSON
	var signupData types.SignupRequest
	if err := c.BodyParser(&signupData); err != nil {
		return RenderHTML(c, auth.ErrOnLogIn(err))
	}

	email := signupData.Email
	password := signupData.Password

	res, err := s.Auth.SignInWithEmailPassword(email, password)

	if err != nil {
		return RenderHTML(c, auth.ErrOnSignUp(err))
	}

	claims := jwt.MapClaims{
		"name":  res.User,
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(k)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72), // Set expiration time
		HTTPOnly: true,                           // Prevents JavaScript access to the cookie -- in any case that we have scripting attacks, the attacker can not steal this information
		Secure:   false,                          // FIXME: Set to true if using HTTPS
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	c.Set("HX-Redirect", "/dashboard")
	return c.SendString("redirecting to /dashboard")
}
