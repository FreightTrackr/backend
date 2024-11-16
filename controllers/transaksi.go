package controllers

import (
	"strconv"

	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var colltransaksi = "transaksi"

func AmbilSemuaTransaksi(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	mconn := utils.SetConnection()
	datatransaksi, datacount, err := utils.GetAllTransaksiWithPagination(mconn, colltransaksi, page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
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

func TambahTransaksi(c *fiber.Ctx) error {
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

	if transaksi.No_Resi == "" || transaksi.Layanan == "" || transaksi.Isi_Kiriman == "" || transaksi.Nama_Pengirim == "" || transaksi.Alamat_Pengirim == "" || transaksi.Kode_Pos_Pengirim == 0 || transaksi.Kota_Asal == "" || transaksi.Nama_Penerima == "" || transaksi.Alamat_Penerima == "" || transaksi.Kode_Pos_Penerima == 0 || transaksi.Kota_Tujuan == "" || transaksi.Berat_Kiriman == 0 || transaksi.Volumetrik == 0 || transaksi.Nilai_Barang == 0 || transaksi.Biaya_Dasar == 0 || transaksi.Biaya_Pajak == 0 || transaksi.Biaya_Asuransi == 0 || transaksi.Total_Biaya == 0 || transaksi.Tanggal_Kirim == primitive.DateTime(0) || transaksi.Tanggal_Terima == primitive.DateTime(0) || transaksi.Tanggal_Tenggat == primitive.DateTime(0) || transaksi.Status == "" || transaksi.Tipe_Cod == "" || transaksi.Status_Cod == "" || transaksi.No_Pend_Kirim == "" || transaksi.No_Pend_Terima == "" || transaksi.Kode_Pelanggan == "" {
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

func HapusTransaksi(c *fiber.Ctx) error {
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
