package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	authservices "github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/auth-services"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/views/components/headers"
)

func SetupHeaderHandlers(server *TASServer) {
	app := server.Server
	header := app.Group("/header")

	header.Get("/action-button", func(c *fiber.Ctx) error {
		user, err := authservices.GetUserFromContext(c)
		if err != nil {
			slog.Info("Could not get user from context" + err.Error())
			return renderHTML(c, headers.ActionButton(nil))
		}

		return renderHTML(c, headers.ActionButton(&user))
	})
}
