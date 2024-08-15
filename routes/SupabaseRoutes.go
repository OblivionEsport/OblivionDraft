package routes

import (
	"oblivion/draft/api"
	"oblivion/draft/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
)

func SupabaseRoutes(app *fiber.App) {
	r := app.Group("/db")

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	r.Use(middleware.SupabaseMiddleware(supabase))

	r.Get("/tournaments", api.GetDBTournaments)
	r.Get("/teams/:tournament_id", api.GetDBTeams)
}