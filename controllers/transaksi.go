package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var colltransaksi = "transaksi"

func FiberAmbilSemuaTransaksiDenganPagination(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid page parameter",
		})
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid limit parameter",
		})
	}
	startDateStr := c.Query("start_date", "")
	endDateStr := c.Query("end_date", "")
	var startDate, endDate time.Time
	no_pend := c.Query("no_pend", "")
	kode_pelanggan := c.Query("kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
	}

	startDate, err = utils.ParseDate(startDateStr, false)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format Start_date tidak valid: " + err.Error(),
		})
	}

	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
	}
	mconn := utils.SetConnection()
	datatransaksi, datacount, err := utils.GetAllTransaksiWithPagination(mconn, colltransaksi, no_pend, kode_pelanggan, page, limit, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datatransaksi == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:     fiber.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datatransaksi,
		Data_Count: &datacount,
		Page:       page,
	})
}

func FiberAmbilSemuaTransaksi(c *fiber.Ctx) error {
	startDateStr := c.Query("start_date", "")
	endDateStr := c.Query("end_date", "")
	var startDate, endDate time.Time
	no_pend := c.Query("no_pend", "")
	kode_pelanggan := c.Query("kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
	}

	startDate, err := utils.ParseDate(startDateStr, false)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format Start_date tidak valid: " + err.Error(),
		})
	}

	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
	}
	mconn := utils.SetConnection()
	datatransaksi, err := utils.GetAllTransaksi(mconn, colltransaksi, no_pend, kode_pelanggan, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datatransaksi == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}

func FiberAmbilTransaksiDenganStatusDelivered(c *fiber.Ctx) error {
	startDateStr := c.Query("start_date", "")
	endDateStr := c.Query("end_date", "")
	var startDate, endDate time.Time
	no_pend := c.Query("no_pend", "")
	kode_pelanggan := c.Query("kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
	}

	startDate, err := utils.ParseDate(startDateStr, false)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format Start_date tidak valid: " + err.Error(),
		})
	}

	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
	}
	mconn := utils.SetConnection()
	datatransaksi, err := utils.GetStatusDeliveredTransaksi(mconn, colltransaksi, no_pend, kode_pelanggan, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datatransaksi == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}

func FiberAmbilTransaksiDenganTipeCOD(c *fiber.Ctx) error {
	startDateStr := c.Query("start_date", "")
	endDateStr := c.Query("end_date", "")
	var startDate, endDate time.Time
	no_pend := c.Query("no_pend", "")
	kode_pelanggan := c.Query("kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
	}

	startDate, err := utils.ParseDate(startDateStr, false)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format Start_date tidak valid: " + err.Error(),
		})
	}

	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
	}
	mconn := utils.SetConnection()
	datatransaksi, err := utils.GetTipeCodTransaksi(mconn, colltransaksi, no_pend, kode_pelanggan, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datatransaksi == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}

func FiberTambahTransaksi(c *fiber.Ctx) error {
	// Peringatan, kode ini belom selesai, no_resi dan id_history belom dibuat generate otomatis

	mconn := utils.SetConnection()
	var transaksi models.Transaksi

	err := c.BodyParser(&transaksi)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Error parsing application/json: " + err.Error(),
		})
	}

	if transaksi.No_Resi == "" || transaksi.Layanan == "" || transaksi.Isi_Kiriman == "" || transaksi.Nama_Pengirim == "" || transaksi.Alamat_Pengirim == "" || transaksi.Kode_Pos_Pengirim == 0 || transaksi.Kota_Asal == "" || transaksi.Nama_Penerima == "" || transaksi.Alamat_Penerima == "" || transaksi.Kode_Pos_Penerima == 0 || transaksi.Kota_Tujuan == "" || transaksi.Berat_Kiriman == 0 || transaksi.Volumetrik == 0 || transaksi.Nilai_Barang == 0 || transaksi.Biaya_Dasar == 0 || transaksi.Biaya_Pajak == 0 || transaksi.Biaya_Asuransi == 0 || transaksi.Total_Biaya == 0 || transaksi.Tanggal_Kirim == primitive.DateTime(0) || transaksi.Tanggal_Antaran_Pertama == primitive.DateTime(0) || transaksi.Tanggal_Terima == primitive.DateTime(0) || transaksi.Status == "" || transaksi.Tipe_Cod == "" || transaksi.Status_Cod == "" || transaksi.No_Pend_Kirim == "" || transaksi.No_Pend_Terima == "" || transaksi.Kode_Pelanggan == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Field wajib diisi",
		})
	}

	if utils.TransaksiExists(mconn, colltransaksi, transaksi) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "No resi sudah ada",
		})
	}

	if transaksi.Layanan != "Reguler" && transaksi.Layanan != "Cepat" && transaksi.Layanan != "Express" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Layanan tidak tersedia",
		})
	}

	if len(transaksi.Isi_Kiriman) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Isi kiriman tidak boleh lebih dari 55 karakter",
		})
	}

	if len(transaksi.Nama_Pengirim) > 55 || len(transaksi.Nama_Penerima) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Nama tidak boleh lebih dari 55 karakter",
		})
	}

	if len(strconv.Itoa(transaksi.Kode_Pos_Pengirim)) > 5 || len(strconv.Itoa(transaksi.Kode_Pos_Penerima)) > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Kode pos tidak boleh lebih dari 5 karakter",
		})
	}

	transaksi.Biaya_Pajak = int(0.11 * float64(transaksi.Biaya_Dasar))
	transaksi.Biaya_Asuransi = int(0.005 * float64(transaksi.Nilai_Barang))
	transaksi.Total_Biaya = transaksi.Nilai_Barang + transaksi.Biaya_Dasar + transaksi.Biaya_Pajak + transaksi.Biaya_Asuransi

	if transaksi.Status != "delivered" && transaksi.Status != "canceled" && transaksi.Status != "returned" && transaksi.Status != "inWarehouse" && transaksi.Status != "inViechle" && transaksi.Status != "failed" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Status tidak valid",
		})
	}

	if transaksi.Tipe_Cod != "noCod" && transaksi.Tipe_Cod != "cod" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Tipe cod tidak tersedia",
		})
	}

	if transaksi.Status_Cod != "paid" && transaksi.Status_Cod != "unPaid" && transaksi.Status_Cod != "onProcess" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Tipe cod tidak tersedia",
		})
	}

	if !utils.KantorExists(mconn, collkantor, models.Kantor{No_Pend: transaksi.No_Pend_Kirim}) || !utils.KantorExists(mconn, collkantor, models.Kantor{No_Pend: transaksi.No_Pend_Terima}) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Kantor tidak ditemukan",
		})
	}

	if !utils.PelangganExists(mconn, collpelanggan, models.Pelanggan{Kode_Pelanggan: transaksi.Kode_Pelanggan}) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Pelanggan tidak ditemukan",
		})
	}

	if !utils.UsernameExists(mconn, collusers, models.Users{Username: transaksi.Created_By.Username}) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "User tidak ditemukan",
		})
	}

	if utils.HistoryExists(mconn, collhistory, models.History{ID_History: transaksi.ID_History}) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "ID Histori sudah ada",
		})
	}

	utils.InsertTransaksi(mconn, colltransaksi, transaksi)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil insert data",
	})
}

func FiberHapusTransaksi(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var transaksi models.Transaksi
	no_resi := c.Query("no_resi")
	transaksi.No_Resi = no_resi
	if !utils.TransaksiExists(mconn, colltransaksi, transaksi) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Transaksi tidak ditemukan",
		})
	}
	utils.DeleteTransaksi(mconn, colltransaksi, transaksi)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil hapus transaksi",
	})
}

func FiberTesting(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid limit parameter",
		})
	}
	startDateStr := c.Query("start_date", "")
	endDateStr := c.Query("end_date", "")
	var startDate, endDate time.Time

	if startDateStr == "" || endDateStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
	}

	startDate, err = utils.ParseDate(startDateStr, false)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format Start_date tidak valid: " + err.Error(),
		})
	}

	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
	}
	mconn := utils.SetConnection()
	datatransaksi, err := utils.GetTransaksiTesting(mconn, colltransaksi, limit, startDate, endDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datatransaksi == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}

func StdAmbilSemuaTransaksiDenganPagination(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	page, err := strconv.Atoi(utils.GetUrlQuery(r, "page", "1"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Invalid page parameter",
		})
		return
	}
	limit, err := strconv.Atoi(utils.GetUrlQuery(r, "limit", "10"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Invalid page parameter",
		})
		return
	}
	startDateStr := utils.GetUrlQuery(r, "start_date", "")
	endDateStr := utils.GetUrlQuery(r, "end_date", "")
	no_pend := utils.GetUrlQuery(r, "no_pend", "")
	kode_pelanggan := utils.GetUrlQuery(r, "kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
		return
	}
	var startDate, endDate time.Time
	startDate, err = utils.ParseDate(startDateStr, false)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format start_date tidak valid: " + err.Error(),
		})
		return
	}
	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
		return
	}
	datatransaksi, datacount, err := utils.GetAllTransaksiWithPagination(mconn, colltransaksi, no_pend, kode_pelanggan, page, limit, startDate, endDate)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
		return
	}
	if datatransaksi == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Pesan{
			Status:  http.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:     http.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datatransaksi,
		Data_Count: &datacount,
		Page:       page,
	})
}

func StdAmbilSemuaTransaksi(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	startDateStr := utils.GetUrlQuery(r, "start_date", "")
	endDateStr := utils.GetUrlQuery(r, "end_date", "")
	no_pend := utils.GetUrlQuery(r, "no_pend", "")
	kode_pelanggan := utils.GetUrlQuery(r, "kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
		return
	}
	var startDate, endDate time.Time
	startDate, err := utils.ParseDate(startDateStr, false)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format start_date tidak valid: " + err.Error(),
		})
		return
	}
	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
		return
	}
	datatransaksi, err := utils.GetAllTransaksi(mconn, colltransaksi, no_pend, kode_pelanggan, startDate, endDate)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
		return
	}
	if datatransaksi == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Pesan{
			Status:  http.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:  http.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}

func StdAmbilTransaksiDenganStatusDelivered(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	startDateStr := utils.GetUrlQuery(r, "start_date", "")
	endDateStr := utils.GetUrlQuery(r, "end_date", "")
	no_pend := utils.GetUrlQuery(r, "no_pend", "")
	kode_pelanggan := utils.GetUrlQuery(r, "kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
		return
	}
	var startDate, endDate time.Time
	startDate, err := utils.ParseDate(startDateStr, false)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format start_date tidak valid: " + err.Error(),
		})
		return
	}
	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
		return
	}
	datatransaksi, err := utils.GetStatusDeliveredTransaksi(mconn, colltransaksi, no_pend, kode_pelanggan, startDate, endDate)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
		return
	}
	if datatransaksi == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Pesan{
			Status:  http.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:  http.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}

func StdAmbilTransaksiDenganTipeCOD(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	startDateStr := utils.GetUrlQuery(r, "start_date", "")
	endDateStr := utils.GetUrlQuery(r, "end_date", "")
	no_pend := utils.GetUrlQuery(r, "no_pend", "")
	kode_pelanggan := utils.GetUrlQuery(r, "kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
		return
	}
	var startDate, endDate time.Time
	startDate, err := utils.ParseDate(startDateStr, false)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format start_date tidak valid: " + err.Error(),
		})
		return
	}
	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
		return
	}
	datatransaksi, err := utils.GetTipeCodTransaksi(mconn, colltransaksi, no_pend, kode_pelanggan, startDate, endDate)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
		return
	}
	if datatransaksi == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Pesan{
			Status:  http.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:  http.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}

func StdExportCSV(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	startDateStr := utils.GetUrlQuery(r, "start_date", "")
	endDateStr := utils.GetUrlQuery(r, "end_date", "")
	no_pend := utils.GetUrlQuery(r, "no_pend", "")
	kode_pelanggan := utils.GetUrlQuery(r, "kode_pelanggan", "")

	if startDateStr == "" || endDateStr == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
		return
	}
	var startDate, endDate time.Time
	startDate, err := utils.ParseDate(startDateStr, false)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format start_date tidak valid: " + err.Error(),
		})
		return
	}
	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
		return
	}
	datatransaksi, err := utils.GetAllTransaksi(mconn, colltransaksi, no_pend, kode_pelanggan, startDate, endDate)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
		return
	}
	if datatransaksi == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Pesan{
			Status:  http.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
		return
	}

	// Set CSV response headers
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", `attachment; filename="transaksi.csv"`)

	writer := csv.NewWriter(w)
	defer writer.Flush()

	// Menulis header CSV
	header := []string{
		"id", "no_resi", "layanan", "isi_kiriman", "nama_pengirim", "alamat_pengirim", "kode_pos_pengirim",
		"kota_asal", "nama_penerima", "alamat_penerima", "kode_pos_penerima", "kota_tujuan",
		"berat_kiriman", "volumetrik", "nilai_barang", "biaya_dasar", "biaya_pajak",
		"biaya_asuransi", "total_biaya", "tanggal_kirim", "tanggal_antaran_pertama",
		"tanggal_terima", "status", "tipe_cod", "status_cod", "sla", "aktual_sla",
		"status_sla", "no_pend_kirim", "no_pend_terima", "kode_pelanggan", "created_by", "id_history",
	}
	if err := writer.Write(header); err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, models.Pesan{
			Status:  http.StatusInternalServerError,
			Message: "Error writing CSV header: " + err.Error(),
		})
		return
	}

	// Menulis data transaksi ke CSV
	for _, transaksi := range datatransaksi {
		row := []string{
			transaksi.ID.Hex(),
			transaksi.No_Resi,
			transaksi.Layanan,
			transaksi.Isi_Kiriman,
			transaksi.Nama_Pengirim,
			transaksi.Alamat_Pengirim,
			strconv.Itoa(transaksi.Kode_Pos_Pengirim),
			transaksi.Kota_Asal,
			transaksi.Nama_Penerima,
			transaksi.Alamat_Penerima,
			strconv.Itoa(transaksi.Kode_Pos_Penerima),
			transaksi.Kota_Tujuan,
			fmt.Sprint(transaksi.Berat_Kiriman),
			fmt.Sprint(transaksi.Volumetrik),
			strconv.Itoa(transaksi.Nilai_Barang),
			strconv.Itoa(transaksi.Biaya_Dasar),
			strconv.Itoa(transaksi.Biaya_Pajak),
			strconv.Itoa(transaksi.Biaya_Asuransi),
			strconv.Itoa(transaksi.Total_Biaya),
			transaksi.Tanggal_Kirim.Time().String(),
			transaksi.Tanggal_Antaran_Pertama.Time().String(),
			transaksi.Tanggal_Terima.Time().String(),
			transaksi.Status,
			transaksi.Tipe_Cod,
			transaksi.Status_Cod,
			strconv.Itoa(transaksi.Sla),
			strconv.Itoa(transaksi.Aktual_Sla),
			fmt.Sprint(transaksi.Status_Sla),
			transaksi.No_Pend_Kirim,
			transaksi.No_Pend_Terima,
			transaksi.Kode_Pelanggan,
			transaksi.Created_By.Username,
			transaksi.ID_History,
		}
		if err := writer.Write(row); err != nil {
			utils.WriteJSONResponse(w, http.StatusInternalServerError, models.Pesan{
				Status:  http.StatusInternalServerError,
				Message: "Error writing CSV row: " + err.Error(),
			})
			return
		}
	}
}

func StdTesting(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	limit, err := strconv.Atoi(utils.GetUrlQuery(r, "limit", "10"))
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Invalid page parameter",
		})
		return
	}
	startDateStr := utils.GetUrlQuery(r, "start_date", "")
	endDateStr := utils.GetUrlQuery(r, "end_date", "")

	if startDateStr == "" || endDateStr == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Masukkan parameter tanggal",
		})
		return
	}
	var startDate, endDate time.Time
	startDate, err = utils.ParseDate(startDateStr, false)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format start_date tidak valid: " + err.Error(),
		})
		return
	}
	endDate, err = utils.ParseDate(endDateStr, true)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Format end_date tidak valid: " + err.Error(),
		})
		return
	}
	datatransaksi, err := utils.GetTransaksiTesting(mconn, colltransaksi, limit, startDate, endDate)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
		return
	}
	if datatransaksi == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Pesan{
			Status:  http.StatusNotFound,
			Message: "Data transaksi tidak ditemukan",
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:  http.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datatransaksi,
	})
}
