package routes

import (
	"net/http"

	"github.com/FreightTrackr/backend/controllers"
	"github.com/FreightTrackr/backend/middleware"
)

func StdRoute(router *http.ServeMux) {
	router.HandleFunc("POST /register", controllers.StdRegister)
	router.HandleFunc("POST /login", controllers.StdLogin)
	router.HandleFunc("GET /lacak", controllers.StdAmbilHistory)
	router.Handle("GET /users", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaUser)))
	router.Handle("PUT /users", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdEditUser)))
	router.Handle("DELETE /users", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdHapusUser)))
	router.Handle("GET /session", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdSession)))
	router.Handle("GET /kantor", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaKantor)))
	router.Handle("GET /pelanggan", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaPelanggan)))
	router.Handle("GET /transaksi", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaTransaksiDenganPagination)))
	router.Handle("GET /semuatransaksi", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilSemuaTransaksi)))
	router.Handle("GET /transaksidelivered", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilTransaksiDenganStatusDelivered)))
	router.Handle("GET /transaksicod", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdAmbilTransaksiDenganTipeCOD)))
	router.Handle("GET /export-csv", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdExportCSV)))
	router.Handle("GET /testing", middleware.IsAuthenticated(http.HandlerFunc(controllers.StdTesting)))
}
