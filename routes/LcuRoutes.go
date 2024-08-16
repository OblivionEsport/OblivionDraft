package routes

import (
	"fmt"
	"oblivion/draft/api"
	"oblivion/draft/middleware"
	"oblivion/draft/models"
	"time"

	"github.com/ImOlli/go-lcu/lcu"
	"github.com/gofiber/fiber/v2"
)

func LcuRoutes(app *fiber.App) {
	var g models.LcuGetter
	var err error
	for {
		g, err = models.NewLcuClient()
		if err != nil && !lcu.IsProcessNotFoundError(err) {
			panic(err)
		} else if err == nil {
			break
		}
		fmt.Println("LeagueClient not found, retrying in 2 second")
		time.Sleep(2 * time.Second)
	}

	r := app.Group("/api/draft")
	r.Use(middleware.DraftMiddleware(g))

	//  ----------- Draft Routes ----------- //
	r.Get("/full", api.DraftFull)
	r.Get("/actions", api.DraftActions)
	r.Get("/bans", api.DraftBans)
	r.Get("/timer", api.DraftTimer)

	r.Get("/summoner/info/", api.SummonerInfo)
	r.Get("/summoner/:id", api.DraftSummoner)
}
