package main

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadEnv(".env")

	app := fiber.New()

	app.Use(cors.New(config.CorsAllowAll))

	routes.FiberRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
