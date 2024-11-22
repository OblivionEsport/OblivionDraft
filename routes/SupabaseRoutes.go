package routes

import (
	"log"
	"oblivion/draft/api"
	"oblivion/draft/middleware"
	"os"

	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
	supa_new "github.com/supabase-community/supabase-go"
)

func SupabaseRoutes(app *fiber.App) {
	r := app.Group("/api/db")

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseUrl == "" || supabaseKey == "" {
		log.Println("Supabase URL or Key is missing, skipping Supabase routes")
		return
	}

	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	supabase_new, err := supa_new.NewClient(supabaseUrl, supabaseKey, &supa_new.ClientOptions{})

	if err != nil {
		log.Println("Error creating Supabase client:", err)
	}

	r.Use(middleware.SupabaseMiddleware(supabase, supabase_new))

	r.Get("/tournaments", api.GetDBTournaments)
	r.Get("/teams/:tournament_id", api.GetDBTeams)
	r.Get("/ewc/stats", api.GetDBStats)
}
