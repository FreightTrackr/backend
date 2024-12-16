package routes

import (
	"net/http"

	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
)

func StdRoute(router *http.ServeMux) {
	router.Handle("POST /std/register", config.IsAuthenticated(http.HandlerFunc(controllers.StdRegister)))
	router.HandleFunc("POST /std/login", controllers.StdLogin)
	router.Handle("GET /std/users", config.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaUser)))
	router.Handle("GET /std/session", config.IsAuthenticated(http.HandlerFunc(controllers.StdSession)))
}
