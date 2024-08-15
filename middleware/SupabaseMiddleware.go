package middleware

import (
	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
)

func SupabaseMiddleware(supabase *supa.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("supabase", supabase)
		return c.Next()
	}
}
