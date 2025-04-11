package api

import (
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
)

func GetDBTeams(c *fiber.Ctx) error {
	supabase := c.Locals("supabase").(*supa.Client)
	tournamentID := c.Params("tournament_id")

	var results []models.DBFetchTeam
	err := supabase.DB.From("part_of").Select("team_id(*)").Eq("tournament_id", tournamentID).Eq("status", "accepted").Execute(&results)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var returnedTeams []models.DBTeam = make([]models.DBTeam, len(results))

	// parse DB fetch team to DB team
	for i := range results {
		returnedTeams[i].ID = results[i].TeamID.ID
		returnedTeams[i].Name = results[i].TeamID.Name
		returnedTeams[i].Tag = results[i].TeamID.Tag
		returnedTeams[i].LogoUrl = results[i].TeamID.LogoUrl
	}

	return c.JSON(returnedTeams)
}
