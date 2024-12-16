package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func FiberRoute(page *fiber.App) {
	page.Post("/fiber/register", controllers.FiberRegister)
	page.Post("/fiber/login", controllers.FiberLogin)
	page.Use(config.JwtMiddleware())
	page.Get("/fiber/users", controllers.FiberAmbilSemuaUser)
	page.Put("/fiber/users", controllers.FiberEditUser)
	page.Delete("/fiber/users", controllers.FiberHapusUser)
	page.Get("/fiber/session", controllers.FiberSession)
	page.Get("/fiber/kantor", controllers.FiberAmbilSemuaKantor)
	page.Post("/fiber/kantor", controllers.FiberTambahKantor)
	page.Delete("/fiber/kantor", controllers.FiberHapusKantor)
	page.Get("/fiber/pelanggan", controllers.FiberAmbilSemuaPelanggan)
	page.Post("/fiber/pelanggan", controllers.FiberTambahPelanggan)
	page.Delete("/fiber/pelanggan", controllers.FiberHapusPelanggan)
	page.Get("/fiber/transaksi", controllers.FiberAmbilSemuaTransaksiDenganPagination)
	page.Get("/fiber/semuatransaksi", controllers.FiberAmbilSemuaTransaksi)
	page.Get("/fiber/transaksidelivered", controllers.FiberAmbilTransaksiDenganStatusDelivered)
	page.Post("/fiber/transaksi", controllers.FiberTambahTransaksi)
	page.Delete("/fiber/transaksi", controllers.FiberHapusTransaksi)
}
