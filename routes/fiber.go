package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func FiberRoute(page *fiber.App) {
	page.Post("/register", controllers.FiberRegister)
	page.Post("/login", controllers.FiberLogin)
	page.Use(config.JwtMiddleware())
	page.Get("/users", controllers.FiberAmbilSemuaUser)
	page.Put("/users", controllers.FiberEditUser)
	page.Delete("/users", controllers.FiberHapusUser)
	page.Get("/session", controllers.FiberSession)
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
