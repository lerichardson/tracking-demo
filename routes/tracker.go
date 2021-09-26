package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Tracker(app *fiber.App) {
	app.Get("/js/tracker.js", func(c *fiber.Ctx) error {
		return c.SendFile("./public/tracker.js")
	})
	return
}
