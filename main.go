package main

import (
	"oblivion/draft/routes"
	"oblivion/draft/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func serve() {

	app := fiber.New()
	// use Logger middleware provided by Fiber
	//app.Use(logger.New())
	app.Get("/metrics", monitor.New())

	routes.StaticRoutes(app)
	routes.AdminRoutes(app)
	routes.DraftRoutes(app)

	app.Listen(":80")
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "setup" {
		utils.Setup()
	} else if len(os.Args) > 1 && os.Args[1] == "update" {
		utils.UpdateOverlay()
	} else {
		utils.CheckSetup()
		serve()
	}
}
