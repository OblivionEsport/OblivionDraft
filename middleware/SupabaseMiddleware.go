package middleware

import (
	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
	supa_new "github.com/supabase-community/supabase-go"
)

func SupabaseMiddleware(supabase *supa.Client, supabase_new *supa_new.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("supabase", supabase)
		c.Locals("supabase_new", supabase_new)
		return c.Next()
	}
}
