package routes

import (
	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(page *fiber.App) {
	page.Post("/register", controllers.Register)
	page.Post("/login", controllers.Login)
	page.Use(config.JwtMiddleware())
	page.Get("/users", controllers.AmbilSemuaUser)
	page.Put("/users", controllers.EditUser)
	page.Delete("/users", controllers.HapusUser)
	page.Get("/session", controllers.Session)
}
