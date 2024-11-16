package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func KantorRoute(page *fiber.App) {
	page.Use(config.JwtMiddleware())
	page.Get("/kantor", controllers.AmbilSemuaKantor)
	page.Post("/kantor", controllers.TambahKantor)
	page.Delete("/kantor", controllers.HapusKantor)
}
