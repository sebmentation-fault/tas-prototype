package authservices

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/sebmentation-fault/tas-prototype/gohtmx-prototype/pkg/db"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Check for authentication credentials
	user, err := GetUserFromContext(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("Unauthorized or malformed jwt")
	}

	c.Locals("user_id", user.GetID())
	c.Locals("user_type", user.GetUserType())
	c.Locals("user_name", user.GetName())

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
