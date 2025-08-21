package api

import (
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
)

func GetFearless(c *fiber.Ctx) error {
	fearless, err := models.ReadTeamsFile()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fearless.Fearless)
}
