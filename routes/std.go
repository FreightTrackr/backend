package routes

import (
	"net/http"

	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/controllers"
)

func StdRoute(router *http.ServeMux) {
	router.Handle("/std/register", config.IsAuthenticated(http.HandlerFunc(controllers.StdRegister)))
	router.HandleFunc("POST /std/login", controllers.StdLogin)
}
