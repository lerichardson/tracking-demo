package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	app.Get("/demo", func(c *fiber.Ctx) error {
		return c.SendFile("./public/demo.html")
	})
	app.Get("/js/tracker.js", func(c *fiber.Ctx) error {
		return c.SendFile("./public/tracker.js")
	})
	app.Post("/trackerData", func(c *fiber.Ctx) error {
		return c.Send(c.Body())
	})
	app.Listen(":3000")
}
