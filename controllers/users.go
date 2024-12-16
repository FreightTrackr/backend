package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var collusers = "users"

func FiberRegister(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var user models.Users

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Error parsing application/json: " + err.Error(),
		})
	}

	if user.Username == "" || user.Password == "" || user.Nama == "" || user.No_Telp == "" || user.Email == "" || user.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Field wajib diisi",
		})
	}

	if strings.Contains(user.Username, " ") || strings.Contains(user.Password, " ") {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Field tidak boleh mengandung spasi",
		})
	}

	if len(user.Username) > 20 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Username tidak boleh lebih dari 20 karakter",
		})
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Password harus antara 8 sampai 20 karakter",
		})
	}

	if len(user.Nama) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Nama tidak boleh lebih dari 55 karakter",
		})
	}

	if user.Role != "admin" && user.Role != "kantor" && user.Role != "pelanggan" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Role tidak tersedia",
		})
	}

	if utils.UsernameExists(mconn, collusers, user) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Username telah dipakai",
		})
	}

	hash, hashErr := helpers.HashPassword(user.Password)
	if hashErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Gagal hash password: " + hashErr.Error(),
		})
	}
	user.Password = hash

	utils.InsertUser(mconn, collusers, user)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil register",
	})
}

func FiberLogin(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var user models.Users

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Error parsing application/json: " + err.Error(),
		})
	}

	if !utils.UsernameExists(mconn, collusers, user) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Akun tidak ditemukan",
		})
	}

	if !utils.IsPasswordValid(mconn, collusers, user) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Password salah",
		})
	}

	token, err := utils.SignedJWT(mconn, collusers, user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Pesan{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil login",
		Token:   token,
	})
}

func FiberAmbilSemuaUser(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	mconn := utils.SetConnection()
	datauser, datacount, err := utils.GetAllUserWithPagination(mconn, collusers, page, limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
	}
	if datauser == nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Pesan{
			Status:  fiber.StatusNotFound,
			Message: "Data user tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:     fiber.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datauser,
		Data_Count: &datacount,
		Page:       page,
	})
}

func FiberEditUser(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var user models.Users

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Error parsing application/json: " + err.Error(),
		})
	}

	if user.Username == "" || user.Nama == "" || user.No_Telp == "" || user.Email == "" || user.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Field wajib diisi",
		})
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Password harus antara 8 sampai 20 karakter",
		})
	}

	if len(user.Nama) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Nama tidak boleh lebih dari 55 karakter",
		})
	}

	if !utils.UsernameExists(mconn, collusers, user) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Username tidak ditemukan",
		})
	}

	if user.Password != "" {
		hash, hashErr := helpers.HashPassword(user.Password)
		if hashErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
				Status:  fiber.StatusBadRequest,
				Message: "Gagal hash password: " + hashErr.Error(),
			})
		}
		user.Password = hash
	} else {
		datauser := utils.FindUser(mconn, collusers, user)
		user.Password = datauser.Password
	}

	utils.InsertUser(mconn, collusers, user)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil update",
	})
}

func FiberHapusUser(c *fiber.Ctx) error {
	mconn := utils.SetConnection()
	var user models.Users
	username := c.Query("username")
	user.Username = username
	if !utils.UsernameExists(mconn, collusers, user) {
		return c.Status(fiber.StatusBadRequest).JSON(models.Pesan{
			Status:  fiber.StatusBadRequest,
			Message: "Akun tidak ditemukan",
		})
	}
	utils.DeleteUser(mconn, collusers, user)
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil hapus user",
	})
}

func FiberSession(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var session models.Users

	session.Username = claims["username"].(string)
	session.Nama = claims["nama"].(string)
	session.No_Telp = claims["no_telp"].(string)
	session.Email = claims["email"].(string)
	session.Role = claims["role"].(string)
	session.No_Pend = claims["no_pend"].(string)
	session.Kode_Pelanggan = claims["kode_pengguna"].(string)

	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berikut data session anda",
		Data:    session,
	})
}

func StdRegister(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	var user models.Users

	utils.ParseBody(w, r, &user)

	if user.Username == "" || user.Password == "" || user.Nama == "" || user.No_Telp == "" || user.Email == "" || user.Role == "" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Field wajib diisi",
		})
		return
	}

	if strings.Contains(user.Username, " ") || strings.Contains(user.Password, " ") {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Field tidak boleh mengandung spasi",
		})
		return
	}

	if len(user.Username) > 20 {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Username tidak boleh lebih dari 20 karakter",
		})
		return
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Password harus antara 8 sampai 20 karakter",
		})
		return
	}

	if len(user.Nama) > 55 {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Nama tidak boleh lebih dari 55 karakter",
		})
		return
	}

	if user.Role != "admin" && user.Role != "kantor" && user.Role != "pelanggan" {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Role tidak tersedia",
		})
		return
	}

	if utils.UsernameExists(mconn, collusers, user) {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Username telah dipakai",
		})
		return
	}

	hash, hashErr := helpers.HashPassword(user.Password)
	if hashErr != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Gagal hash password: " + hashErr.Error(),
		})
		return
	}
	user.Password = hash

	utils.InsertUser(mconn, collusers, user)
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:  http.StatusOK,
		Message: "Berhasil register",
	})
}

func StdLogin(w http.ResponseWriter, r *http.Request) {
	mconn := utils.SetConnection()
	var user models.Users

	utils.ParseBody(w, r, &user)

	if !utils.UsernameExists(mconn, collusers, user) {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Akun tidak ditemukan",
		})
		return
	}

	if !utils.IsPasswordValid(mconn, collusers, user) {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Password salah",
		})
		return
	}

	token, err := utils.SignedJWT(mconn, collusers, user)

	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, models.Pesan{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:  http.StatusOK,
		Message: "Berhasil login",
		Token:   token,
	})
}

func StdAmbilSemuaUser(w http.ResponseWriter, r *http.Request) {
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
	datauser, datacount, err := utils.GetAllUserWithPagination(mconn, collusers, page, limit)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "GetAllDoc error: " + err.Error(),
		})
		return
	}
	if datauser == nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Pesan{
			Status:  http.StatusNotFound,
			Message: "Data user tidak ditemukan",
		})
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, models.Pesan{
		Status:     http.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datauser,
		Data_Count: &datacount,
		Page:       page,
	})
}
