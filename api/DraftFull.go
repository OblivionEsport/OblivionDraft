package api

import (
	"oblivion/draft/models"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func DraftFull(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.LcuGetter)
	s, err := utils.GetDraft(g)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":   "Draft not found",
			"message": err.Error(),
		})
	}
	return c.JSON(s)
}
