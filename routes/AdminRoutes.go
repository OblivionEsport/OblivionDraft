package routes

import (
	"oblivion/draft/api"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	r := app.Group("/admin")
	teams := r.Group("/teams")

	teams.Get("/full", api.TeamsFull)
	teams.Get("/selected", api.GetTeamsSelected)
	teams.Get("/delete/:name", api.TeamsDelete)

	teams.Post("/add", api.TeamsAdd)
	teams.Post("/selected", api.SetTeamsSelected)

	r.Get("/match/id", api.MatchID)
	r.Post("/match/id", api.MatchID)
}
