package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func FiberRoute(page *fiber.App) {
	page.Post("/register", controllers.Register)
	page.Post("/login", controllers.Login)
	page.Use(config.JwtMiddleware())
	page.Get("/users", controllers.AmbilSemuaUser)
	page.Put("/users", controllers.EditUser)
	page.Delete("/users", controllers.HapusUser)
	page.Get("/session", controllers.Session)
	page.Get("/kantor", controllers.AmbilSemuaKantor)
	page.Post("/kantor", controllers.TambahKantor)
	page.Delete("/kantor", controllers.HapusKantor)
	page.Get("/pelanggan", controllers.AmbilSemuaPelanggan)
	page.Post("/pelanggan", controllers.TambahPelanggan)
	page.Delete("/pelanggan", controllers.HapusPelanggan)
	page.Get("/transaksi", controllers.AmbilSemuaTransaksiDenganPagination)
	page.Get("/semuatransaksi", controllers.AmbilSemuaTransaksi)
	page.Get("/transaksidelivered", controllers.AmbilTransaksiDenganStatusDelivered)
	page.Post("/transaksi", controllers.TambahTransaksi)
	page.Delete("/transaksi", controllers.HapusTransaksi)
}
