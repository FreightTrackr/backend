package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func TransaksiRoute(page *fiber.App) {
	page.Use(config.JwtMiddleware())
	page.Get("/transaksi", controllers.AmbilSemuaTransaksi)
	page.Get("/transaksivisual", controllers.AmbilSemuaTransaksiUntukVisualusasi)
	page.Post("/transaksi", controllers.TambahTransaksi)
	page.Delete("/transaksi", controllers.HapusTransaksi)
}
