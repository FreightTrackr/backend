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
	router.Handle("GET /std/kantor", config.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaKantor)))
	router.Handle("GET /std/pelanggan", config.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaPelanggan)))
	router.Handle("GET /std/transaksi", config.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaTransaksiDenganPagination)))
	router.Handle("GET /std/semuatransaksi", config.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaTransaksi)))
}
