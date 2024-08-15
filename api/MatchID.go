package api

import (
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func MatchID(c *fiber.Ctx) error {
	// if GET, return matchID from file, else update it from matchID params

	// if POST, update matchID, else return it
	if c.Method() == "POST" {
		var matchID string
		body := c.Body()
		matchID = string(body)
		utils.SetMatchID(matchID)
		return c.SendStatus(201)
	}

	// if GET, return matchID
	matchID, err := utils.GetMatchID()
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendString(matchID)
}
