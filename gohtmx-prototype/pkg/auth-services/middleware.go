package authservices

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/db"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Check for authentication credentials
	tokenString := c.Cookies("jwt")

	token, err := ValidateJWTToken(tokenString)

	if err != nil {
		slog.Info("Problem parsing the token: ", err)
		c.Status(fiber.StatusUnprocessableEntity)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("Problem encountered while parsing the token")
	}

	if !token.Valid {
		slog.Info("User is using an invalid token")
		c.Status(fiber.StatusUnauthorized)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("Invalid token")
	}

	id, userType, err := GetAccountIdAndTypeFromToken(token)
	if err != nil {
		slog.Info("Token could not be parsed")
		c.Status(fiber.StatusUnprocessableEntity)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("Invalid token")
	}

	c.Locals("user_id", id)
	c.Locals("user_type", userType)

	// If credentials are valid, proceed to the next middleware or route
	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	// get the user role/type
	userType := c.Locals("user_type").(int)

	// if not admin, error
	if userType != int(db.AdminUser) {
		slog.Info("A user has tried to reach the admin site")
		c.Status(fiber.StatusForbidden)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("Permission denied. You are not an admin.")
	}

	return c.Next()
}

func CelebrityMiddleware(c *fiber.Ctx) error {
	// get the user role/type
	userType := c.Locals("user_type").(int)

	// if not celeb, error
	if userType != int(db.CelebrityUser) {
		slog.Info("A user has tried to reach the admin site")
		c.Status(fiber.StatusForbidden)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("Permission denied. You are not a celebrity.")
	}

	return c.Next()
}
