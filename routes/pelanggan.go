package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func PelangganRoute(page *fiber.App) {
	page.Use(config.JwtMiddleware())
	page.Get("/pelanggan", controllers.AmbilSemuaPelanggan)
	page.Post("/pelanggan", controllers.TambahPelanggan)
	page.Delete("/pelanggan", controllers.HapusPelanggan)
}
