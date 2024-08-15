package api

import (
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
)

func GetDBTournaments(c *fiber.Ctx) error {
	supabase := c.Locals("supabase").(*supa.Client)

	var r []models.DBTournament
	err := supabase.DB.From("Tournaments").Select("*").Execute(&r)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(r)
}
