package api

import (
	"log"
	"oblivion/draft/models"
	"oblivion/draft/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DraftSummoner(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.LcuGetter)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.SendStatus(404)
	}
	s, err := utils.GetDraft(g)
	if err != nil {
		return c.SendStatus(404)
	}
	if len(s.MyTeam) < 5 || len(s.TheirTeam) < 5 {
		return c.SendStatus(405)
	}
	if id < 5 {
		return c.JSON(s.MyTeam[id])
	}
	return c.JSON(s.TheirTeam[id-5])
}
