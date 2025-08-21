package api

import (
	"oblivion/draft/models"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func SummonerInfo(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.LcuGetter)

	championID := c.Query("championID")
	if championID == "" {
		return c.SendStatus(404)
	}
	champName, splashUrl, iconURL := utils.GetChampionInfo(championID)

	summonerID := c.Query("summonerID")
	if summonerID == "" {
		return c.JSON(fiber.Map{"championName": champName, "splashUrl": splashUrl, "iconURL": iconURL})
	}
	name := utils.GetName(g, summonerID)

	return c.JSON(fiber.Map{"summonerName": name, "championName": champName, "splashUrl": splashUrl, "iconURL": iconURL})
}
