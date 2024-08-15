package api

import (
	"encoding/json"
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
)

func MatchHistoryFull(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.ApiGetter)
	matchID := c.Params("id")
	rawStats := g.Getf("/lol/match/v5/matches/%s/timeline", matchID)
	return c.JSON(json.RawMessage(rawStats))
}
