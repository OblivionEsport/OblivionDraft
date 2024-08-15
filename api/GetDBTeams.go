package api

import (
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
)

func GetDBTeams(c *fiber.Ctx) error {
	supabase := c.Locals("supabase").(*supa.Client)
	tournamentID := c.Params("tournament_id")

	var results []models.DBTeam
	err := supabase.DB.From("Teams").Select("*").Eq("tournament_id", tournamentID).Execute(&results)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(results)
}
