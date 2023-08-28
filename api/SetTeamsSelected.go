package api

import (
	"log"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func SetTeamsSelected(c *fiber.Ctx) error {
	selected := new([]string)
	if err := c.BodyParser(selected); err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	err := utils.SetSelected((*selected)[0], (*selected)[1])
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
