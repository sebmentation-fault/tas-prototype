package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/auth"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/layouts"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

// the first two handlers simply render the html

func signUpHandler(c *fiber.Ctx) error {
	return RenderHTML(c, layouts.Base("Sign Up", auth.SignUpTempl()))
}

func logInHandler(c *fiber.Ctx) error {
	return c.Redirect("/404-not-found")
}

// these two handlers are going to respond to the log in/sign up form submition

func signUpSubmittedHandlerWrapper(s *supabase.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return signUpSubmittedHandler(c, s)
	}
}

func signUpSubmittedHandler(c *fiber.Ctx, s *supabase.Client) error {
	fmt.Printf("email: %s, password: %s\n", c.FormValue("email"), c.FormValue("password"))

	email := "sebastian@kjallgren.com" //c.FormValue("email")
	password := "password1"            //c.FormValue("password")

	req := types.SignupRequest{
		Email:    email,
		Password: password,
	}

	res, err := s.Auth.Signup(req)

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
	return c.Redirect("/404-not-found")
}
