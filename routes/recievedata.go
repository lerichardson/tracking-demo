package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RecieveTrackerData(app *fiber.App) {
	app.Post("/trackerData", func(c *fiber.Ctx) error {
		return c.Send(c.Body())
	})
	return
}
