package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lerichardson/tracking-demo/routes"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	routes.Tracker(app)
	routes.RecieveTrackerData(app)
	routes.DemoPage(app)
	app.Listen(":3000")
}
