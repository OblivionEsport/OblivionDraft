package api

import (
	"log"
	"oblivion/draft/models"
	"oblivion/draft/utils"

	"github.com/gofiber/fiber/v2"
)

func DraftActions(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.LcuGetter)
	s, err := utils.GetActions(g)
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	return c.JSON(s)
}
