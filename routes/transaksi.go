package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func TransaksiRoute(page *fiber.App) {
	page.Use(config.JwtMiddleware())
	page.Get("/transaksi", controllers.AmbilSemuaTransaksiDenganPagination)
	page.Get("/semuatransaksi", controllers.AmbilSemuaTransaksi)
	page.Get("/transaksistatusdelivered", controllers.AmbilTransaksiDenganStatusDelivered)
	page.Post("/transaksi", controllers.TambahTransaksi)
	page.Delete("/transaksi", controllers.HapusTransaksi)
}
