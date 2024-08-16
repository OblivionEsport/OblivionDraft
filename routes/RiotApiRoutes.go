package routes

import (
	"oblivion/draft/api"
	"oblivion/draft/middleware"
	"oblivion/draft/models"
	"os"

	"github.com/gofiber/fiber/v2"
)

func RiotApiRoutes(app *fiber.App) {
	apiKey := os.Getenv("API_KEY")
	g, err := models.NewApiClient(apiKey, models.Region("europe"))
	if err != nil {
		panic(err)
	}

	r := app.Group("/api/riot")
	r.Use(middleware.RiotMiddleware(g))

	r.Get("/puuid/:name/:tag", api.APIGetPuuid)
	r.Get("/current-spect-id/:puuid", api.TeamsFull)
	r.Get("/match/:id/timeline", api.MatchHistoryFull)
	r.Get("/match/:id/endgame", api.EndGameStats)

	r.Get("/champ/:championID", api.GetChampInfo)

}
