package api

import (
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteFearless(c *fiber.Ctx) error {
	if err := utils.DeleteFearlessBans(); err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
