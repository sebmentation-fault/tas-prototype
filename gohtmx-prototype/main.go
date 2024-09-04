package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/handlers"
	"github.com/supabase-community/supabase-go"
)

func main() {
	// Set up the app
	app := fiber.New()

	// FIXME: the key (and url) should be fetched from env variables

	// as defined in supabase -> docker -> env -> site_url
	const API_URL = "http://localhost:8000"
	// as defined in supabase -> docker -> env -> anon_key
	const API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ewogICJyb2xlIjogImFub24iLAogICJpc3MiOiAic3VwYWJhc2UiLAogICJpYXQiOiAxNzI1MTQ1MjAwLAogICJleHAiOiAxODgyOTExNjAwCn0.NgGJ_9MoaFbyaAyPdlzRLWyA1QkxhIyc31aR1pJmAM8"
	supabaseClient, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
		return
	}

	// Middleware to cache the favicon, speeding up acces to favicon requests
	// Also corrects the URL
	app.Use(favicon.New(favicon.Config{
		File: "./public/favicon.ico",
		URL:  "/favicon.ico",
	}))

	// On client-leveL code, we use '/static' to reach the assets in 'public/',
	// to access CSS, JS, etc.
	app.Static("/static", "./public/")

	// Set up the handlers
	handlers.SetupHandlers(app, supabaseClient)

	// Start listening
	app.Listen(":8080")
}
