package api

import (
	"oblivion/draft/models"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func SummonerInfo(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.Getter)
	summonerID := c.Query("summonerID")
	if summonerID == "" {
		return c.SendStatus(404)
	}
	name := utils.GetName(g, summonerID)

	championID := c.Query("championID")
	if championID == "" {
		return c.SendStatus(404)
	}
	champName, splashUrl, iconURL := utils.GetChampionInfo(championID)

	return c.JSON(fiber.Map{"summonerName": name, "championName": champName, "splashUrl": splashUrl, "iconURL": iconURL})
}
