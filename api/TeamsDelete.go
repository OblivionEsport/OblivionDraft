package api

import (
	"log"
	"net/url"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func TeamsDelete(c *fiber.Ctx) error {
	nameURL := c.Params("name")
	reset := c.Query("reset")
	if reset == "true" {
		// remove all teams
		err := utils.ResetTeams()
		if err != nil {
			log.Println(err)
			return c.SendStatus(404)
		}
		return c.SendStatus(200)
	}

	name, err := url.QueryUnescape(nameURL)
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	err = utils.DeleteTeam(name)
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
