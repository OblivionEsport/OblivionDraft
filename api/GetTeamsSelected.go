package api

import (
	"log"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTeamsSelected(c *fiber.Ctx) error {
	selected, err := utils.GetSelected()
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	return c.JSON(selected)
}
