package api

import (
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func GetChampInfo(c *fiber.Ctx) error {
	championID := c.Params("championID")
	if championID == "" {
		return c.SendStatus(404)
	}
	champName, splashUrl, iconURL := utils.GetChampionInfo(championID)
	return c.JSON(fiber.Map{"championName": champName, "splashUrl": splashUrl, "iconURL": iconURL})
}
