package main

import (
	"log"
	"oblivion/draft/routes"
	"oblivion/draft/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func serve() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(err)
			code := fiber.StatusInternalServerError
			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(code).SendString("Internal Server Error")
		},
	})
	app.Get("/metrics", monitor.New())
	app.Use(recover.New())

	routes.StaticRoutes(app)
	routes.AdminRoutes(app)
	routes.LcuRoutes(app)
	routes.RiotApiRoutes(app)
	routes.SupabaseRoutes(app)

	app.Listen(":80")
}

func main() {
	utils.ConfigLogger()
	if len(os.Args) > 1 && os.Args[1] == "setup" {
		utils.Setup()
	} else if len(os.Args) > 1 && os.Args[1] == "update" {
		utils.UpdateOverlay()
	} else {
		utils.CheckSetup()
		serve()

	}
}
