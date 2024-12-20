package routes

import (
	"net/http"

	"github.com/FreightTrackr/backend/controllers"
	"github.com/FreightTrackr/backend/middleware"
)

func StdRoute(router *http.ServeMux) {
	router.Handle("POST /std/register", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdRegister)))
	router.HandleFunc("POST /std/login", controllers.StdLogin)
	router.Handle("GET /std/users", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaUser)))
	router.Handle("PUT /std/users", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdEditUser)))
	router.Handle("GET /std/session", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdSession)))
	router.Handle("GET /std/kantor", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaKantor)))
	router.Handle("GET /std/pelanggan", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaPelanggan)))
	router.Handle("GET /std/transaksi", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaTransaksiDenganPagination)))
	router.Handle("GET /std/semuatransaksi", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaTransaksi)))
	router.Handle("GET /std/transaksidelivered", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilTransaksiDenganStatusDelivered)))
	router.Handle("GET /std/getrole", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdGetRole)))
}
