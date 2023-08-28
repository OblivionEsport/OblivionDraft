package main

import (
	"oblivion/draft/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {

	app := fiber.New()
	// use Logger middleware provided by Fiber
	app.Use(logger.New())
	app.Get("/metrics", monitor.New())

	routes.StaticRoutes(app)
	routes.AdminRoutes(app)
	routes.DraftRoutes(app)

	app.Listen(":80")
}
