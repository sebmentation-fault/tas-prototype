package authservices

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/db"
)

// TODO: cache the user so we do not need this many db calls?
func GetUserFromContext(c *fiber.Ctx) (db.User, error) {

	tokenString := c.Cookies("jwt")

	token, err := ValidateJWTToken(tokenString)
	if err != nil {
		slog.Info("Problem parsing the token: ", err)
		return nil, fiber.ErrUnprocessableEntity
	}

	if !token.Valid {
		slog.Info("User is using an invalid token")
		return nil, fiber.ErrUnauthorized
	}

	user, err := GetUserFromToken(token)
	if err != nil {
		slog.Info("Token could not be parsed")
		return nil, fiber.ErrUnprocessableEntity
	}

	return user, nil
}
