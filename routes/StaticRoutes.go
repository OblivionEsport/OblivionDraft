package routes

import "github.com/gofiber/fiber/v2"

func StaticRoutes(app *fiber.App) {
	app.Static("/", "./overlay/")
	app.Static("/admin/", "./admin/")
}
