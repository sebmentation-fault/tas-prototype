package handlers

import (
	"github.com/gofiber/fiber/v2/middleware/monitor"
	authservices "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/auth-services"
)

// All routes in admin are required to have the JWT token thing
func SetupAdminHandlers(server *TASServer) {
	app := server.Server
	admin := app.Group("/admin")

	// require authenticated users past here
	admin.Use(authservices.AuthMiddleware)
	// require they be admin
	admin.Use(authservices.AdminMiddleware)

	admin.Get("/metrics", monitor.New(monitor.Config{
		Title: "Take a Selfie Metrics",
	})) // /admin/metrics
}
