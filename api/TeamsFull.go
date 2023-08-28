package api

import (
	"log"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func TeamsFull(c *fiber.Ctx) error {
	teams, err := utils.GetTeams()
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	return c.JSON(teams)
}
