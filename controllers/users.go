package controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var collusers = "users"

func Register(c *fiber.Ctx) error {
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

func Login(c *fiber.Ctx) error {
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

	datauser := utils.FindUser(mconn, collusers, user)

	claims := jwt.MapClaims{
		"username":      user.Username,
		"nama":          datauser.Nama,
		"no_telp":       datauser.No_Telp,
		"email":         datauser.Email,
		"role":          datauser.Role,
		"no_pend":       datauser.No_Pend,
		"kode_pengguna": datauser.Kode_Pelanggan,
		"exp":           time.Now().Add(time.Hour * 2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	privateKey, err := utils.ReadPrivateKeyFromFile("./keys/private.pem")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Pesan{
			Status:  fiber.StatusInternalServerError,
			Message: "Error loading private key: " + err.Error(),
		})
	}

	t, err := token.SignedString(privateKey)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:  fiber.StatusOK,
		Message: "Berhasil login",
		Token:   t,
	})
}

func AmbilSemuaUser(c *fiber.Ctx) error {
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
	return c.Status(fiber.StatusOK).JSON(models.Pesan{
		Status:     fiber.StatusOK,
		Message:    "Berhasil ambil data",
		Data:       datauser,
		Data_Count: &datacount,
		Page:       page,
	})
}

func EditUser(c *fiber.Ctx) error {
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

func HapusUser(c *fiber.Ctx) error {
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

func Session(c *fiber.Ctx) error {
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
