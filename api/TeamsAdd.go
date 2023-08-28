package api

import (
	"log"
	"oblivion/draft/models"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func TeamsAdd(c *fiber.Ctx) error {
	team := new(models.Team)
	if err := c.BodyParser(team); err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	new, err := utils.AddTeam(team)
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	if new {
		return c.SendStatus(201)
	}
	return c.SendStatus(200)
}
