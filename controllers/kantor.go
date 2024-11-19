package controllers

import (
	"strconv"

	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
)

var collkantor = "kantor"

func AmbilSemuaKantor(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	mconn := utils.SetConnection()
	datakantor, datacount, err := utils.GetAllKantorWithPagination(mconn, collkantor, page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datakantor == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data kantor tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:     fiber.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datakantor,
		Data_Count: &datacount,
		Page:       page,
	})
}

func TambahKantor(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var kantor models.Kantor

	err := c.BodyParser(&kantor)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Error parsing application/json: " + err.Error(),
		})
	}

	if kantor.No_Pend == "" || kantor.Nama_Kantor == "" || kantor.Kota_Kantor == "" || kantor.Kode_Pos_Kantor == 0 || kantor.Alamat_Kantor == "" {
		if kantor.No_Pend_Kcu == "" || kantor.No_Pend_Kc == "" {
			return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
				Status:  fiber.StatusBadRequest,
				Message: "Field wajib diisi",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Field wajib diisi",
		})
	}

	if utils.KantorExists(mconn, collkantor, kantor) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "No pend sudah ada",
		})
	}

	if kantor.Tipe_Kantor != "kcu" && kantor.Tipe_Kantor != "kc" && kantor.Tipe_Kantor != "kcp" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Tipe pelanggan tidak tersedia",
		})
	}

	if len(kantor.Nama_Kantor) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Nama tidak boleh lebih dari 55 karakter",
		})
	}

	if kantor.Region_Kantor < 1 || kantor.Region_Kantor > 6 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Region tidak tersedia",
		})
	}

	if len(strconv.Itoa(kantor.Kode_Pos_Kantor)) > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Kode pos tidak boleh lebih dari 5 karakter",
		})
	}

	if len(kantor.Alamat_Kantor) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Alamat tidak boleh lebih dari 55 karakter",
		})
	}

	utils.InsertKantor(mconn, collkantor, kantor)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil insert data",
	})
}

func HapusKantor(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var kantor models.Kantor
	no_pend := c.Query("no_pend")
	kantor.No_Pend = no_pend
	if !utils.KantorExists(mconn, collkantor, kantor) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Pelanggan tidak ditemukan",
		})
	}
	utils.DeleteKantor(mconn, collkantor, kantor)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil hapus pelanggan",
	})
}
