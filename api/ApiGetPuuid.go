package api

import (
	"encoding/json"
	"oblivion/draft/models"

	"github.com/gofiber/fiber/v2"
)

func APIGetPuuid(c *fiber.Ctx) error {
	g := c.Locals("getter").(models.ApiGetter)
	name := c.Params("name")
	tag := c.Params("tag")
	rawStats := g.Getf("/riot/account/v1/accounts/by-riot-id/%s/%s", name, tag)
	r := models.PuuidResponse{}
	json.Unmarshal([]byte(rawStats), &r)
	return c.SendString(r.Puuid)

}
