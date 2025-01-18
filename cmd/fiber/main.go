package main

import (
	"github.com/FreightTrackr/backend/middleware"
	"github.com/FreightTrackr/backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	middleware.LoadEnv(".env")

	app := fiber.New()

	app.Use(cors.New(middleware.FiberLocalCors))

	routes.FiberRoute(app)

	app.Listen(":3000")
}
