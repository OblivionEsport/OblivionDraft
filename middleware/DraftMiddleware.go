package middleware

import (
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
)

func DraftMiddleware(g models.Getter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("getter", g)
		return c.Next()
	}
}
