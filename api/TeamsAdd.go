package api

import (
	"log"
	"oblivion/draft/models"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func TeamsAdd(c *fiber.Ctx) error {
	many := c.Query("many")
	if many == "true" {
		teams := new([]models.Team)
		if err := c.BodyParser(teams); err != nil {
			log.Println(err)
			return c.SendStatus(404)
		}
		for _, team := range *teams {
			_, err := utils.AddTeam(&team)
			if err != nil {
				log.Println(err)
				return c.SendStatus(404)
			}
		}
		return c.SendStatus(201)
	} else {
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
	}
	return c.SendStatus(408)
}
