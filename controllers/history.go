package controllers

import (
	"net/http"
	"strconv"

	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
)

var collhistory = "history"

func AmbilSemuaHistory(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	mconn := utils.SetConnection()
	datahistory, datacount, err := utils.GetAllHistoryWithPagination(mconn, collhistory, page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datahistory == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data user tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:     fiber.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datahistory,
		Data_Count: &datacount,
		Page:       page,
	})
}

func StdAmbilHistory(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	var history models.History
	datahistory := utils.FindHistory(mconn, collhistory, history)
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:  http.StatusOK,
		Message: "Berhasil ambil data",
		Data:    datahistory,
	})
}
