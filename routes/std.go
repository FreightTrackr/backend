package routes

import (
	"net/http"

	"github.com/FreightTrackr/backend/controllers"
)

func StdRoute(router *http.ServeMux) {
	router.HandleFunc("POST /std/register", controllers.StdRegister)
}
