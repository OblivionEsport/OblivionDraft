package api

import (
	"log"
	"oblivion/draft/models"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func DraftBans(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.LcuGetter)
	s, err := utils.GetDraft(g)
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	teamFile, err := models.ReadTeamsFile()
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	for _, ban := range s.Bans.MyTeamBans {
		if !contains(teamFile.Fearless[0], ban) {
			teamFile.Fearless[0] = append(teamFile.Fearless[0], ban)
			println("Added to fearless:", ban)
		}
	}
	for _, ban := range s.Bans.TheirTeamBans {
		if !contains(teamFile.Fearless[1], ban) {
			teamFile.Fearless[1] = append(teamFile.Fearless[1], ban)
			println("Added to fearless:", ban)
		}
	}

	// save
	if err := models.SaveTeamsFile(teamFile); err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	return c.JSON(s.Bans)
}

// contains returns true if v exists in list.
func contains[T comparable](list []T, v T) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}
