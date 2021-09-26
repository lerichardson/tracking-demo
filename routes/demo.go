package routes

import (
	"github.com/gofiber/fiber/v2"
)

func DemoPage(app *fiber.App) {
	app.Get("/demo", func(c *fiber.Ctx) error {
		return c.SendFile("./public/demo.html")
	})
	return
}

