package routes

import (
	"oblivion/draft/api"
	"oblivion/draft/middleware"
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
)

func DraftRoutes(app *fiber.App) {
	g, err := models.NewGetter()
	if err != nil {
		panic(err)
	}

	r := app.Group("/draft")
	r.Use(middleware.DraftMiddleware(g))

	r.Get("/full", api.DraftFull)
	r.Get("/actions", api.DraftActions)
	r.Get("/bans", api.DraftBans)
	r.Get("/timer", api.DraftTimer)

	r.Get("/summoner/info/", api.SummonerInfo)
	r.Get("/summoner/:id", api.DraftSummoner)
}
