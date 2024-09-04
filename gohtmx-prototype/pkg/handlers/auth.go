package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/auth"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

func authHandler(c *fiber.Ctx) error {
	return RenderHTML(c, layouts.Base("Authenticate", auth.SignUpTempl()))
}

// the first two handlers simply render the html with the base layout

func signUpHandler(c *fiber.Ctx) error {
	return RenderHTML(c, auth.SignUpTempl())
}

func logInHandler(c *fiber.Ctx) error {
	return RenderHTML(c, auth.LogInTempl())
}

// these two handlers are going to respond to the log in/sign up form submition

func signUpSubmittedHandlerWrapper(s *supabase.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return signUpSubmittedHandler(c, s)
	}
}

func signUpSubmittedHandler(c *fiber.Ctx, s *supabase.Client) error {
	// Parse the request body as JSON
	var signupData types.SignupRequest
	if err := c.BodyParser(&signupData); err != nil {
		return RenderHTML(c, auth.ErrOnSignUp(err))
	}

	res, err := s.Auth.Signup(signupData)

	if err != nil {
		return RenderHTML(c, auth.ErrOnSignUp(err))
	}

	var _ = res

	// user is signed up, we should automagically log them in to now
	// res, err := s.Auth.SignInWithEmailPassword(email, password)
	//
	// if err != nil {
	// 	// FIXME: update the HTML to say log in did not work
	// 	return c.Redirect("/404-not-found")
	// }

	// TODO: add the cookies or whatever
	c.Set("HX-Redirect", "/dashboard")
	return c.SendString("redirecting to /dashboard")
}

func logInSubmittedHandlerWrapper(s *supabase.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return logInSubmittedHandler(c, s)
	}
}

func logInSubmittedHandler(c *fiber.Ctx, s *supabase.Client) error {
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

	// user is logged in now, cool!
	var _ = res

	// TODO: add the cookies or whatever
	c.Set("HX-Redirect", "/dashboard")
	return c.SendString("redirecting to /dashboard")
}
