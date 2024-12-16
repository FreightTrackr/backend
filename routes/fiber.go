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
	page.Get("/kantor", controllers.FiberAmbilSemuaKantor)
	page.Post("/kantor", controllers.FiberTambahKantor)
	page.Delete("/kantor", controllers.FiberHapusKantor)
	page.Get("/pelanggan", controllers.FiberAmbilSemuaPelanggan)
	page.Post("/pelanggan", controllers.FiberTambahPelanggan)
	page.Delete("/pelanggan", controllers.FiberHapusPelanggan)
	page.Get("/transaksi", controllers.AmbilSemuaTransaksiDenganPagination)
	page.Get("/semuatransaksi", controllers.AmbilSemuaTransaksi)
	page.Get("/transaksidelivered", controllers.AmbilTransaksiDenganStatusDelivered)
	page.Post("/transaksi", controllers.TambahTransaksi)
	page.Delete("/transaksi", controllers.HapusTransaksi)
}
