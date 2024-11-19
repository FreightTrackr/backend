package controllers

import (
	"strconv"

	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
)

var collpelanggan = "pelanggan"

func AmbilSemuaPelanggan(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	mconn := utils.SetConnection()
	datapelanggan, datacount, err := utils.GetAllPelangganWithPagination(mconn, collpelanggan, page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datapelanggan == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data pelanggan tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:     fiber.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datapelanggan,
		Data_Count: &datacount,
		Page:       page,
	})
}

func AmbilSemuaPelangganFilter(c *fiber.Ctx) error {
	// Get pagination parameters from query string (or set defaults)
	page, _ := strconv.Atoi(c.Query("page", "1"))              // Default to page 1
	limit, _ := strconv.Atoi(c.Query("limit", "10"))           // Default to 10 items per page
	tipe_pelanggan := c.Query("tipe_pelanggan", "Marketplace") // Default to 10 items per page

	// Establish a connection
	mconn := utils.SetConnection()

	// Fetch paginated users data
	datapelanggan, datacount, err := utils.GetAllPelangganByFilterWithPagination(mconn, collpelanggan, page, limit, tipe_pelanggan)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}

	// Return the filtered and extracted data along with the total number of documents
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:     fiber.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datapelanggan,
		Data_Count: &datacount,
		Page:       page,
	})
}

func TambahPelanggan(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var pelanggan models.Pelanggan

	err := c.BodyParser(&pelanggan)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Error parsing application/json: " + err.Error(),
		})
	}

	if pelanggan.Kode_Pelanggan == "" || pelanggan.Tipe_Pelanggan == "" || pelanggan.Nama_Pelanggan == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Field wajib diisi",
		})
	}

	if utils.PelangganExists(mconn, collpelanggan, pelanggan) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Kode pelanggan sudah ada",
		})
	}

	if pelanggan.Tipe_Pelanggan != "Retail" && pelanggan.Tipe_Pelanggan != "Corporate" && pelanggan.Tipe_Pelanggan != "Marketplace" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Tipe pelanggan tidak tersedia",
		})
	}

	if len(pelanggan.Nama_Pelanggan) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Nama tidak boleh lebih dari 55 karakter",
		})
	}

	utils.InsertPelanggan(mconn, collpelanggan, pelanggan)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil insert data",
	})
}

func HapusPelanggan(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var pelanggan models.Pelanggan
	kode_pelanggan := c.Query("kode_pelanggan")
	pelanggan.Kode_Pelanggan = kode_pelanggan
	if !utils.PelangganExists(mconn, collpelanggan, pelanggan) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Pelanggan tidak ditemukan",
		})
	}
	utils.DeletePelanggan(mconn, collpelanggan, pelanggan)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil hapus pelanggan",
	})
}
