package api

import (
	"encoding/json"
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/postgrest-go"
	supa "github.com/supabase-community/supabase-go"
)

func GetDBStats(c *fiber.Ctx) error {
	supabase := c.Locals("supabase_new").(*supa.Client)

	var results []models.DBStatsEWCWithTeams
	desc := postgrest.OrderOpts{Ascending: false}
	data, _, err := supabase.From("stats_ewc").Select("match_id(team_one(logo_url), team_two(logo_url)), stats", "exact", false).Order("match_id", &desc).Execute()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = json.Unmarshal(data, &results)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(results)
}
