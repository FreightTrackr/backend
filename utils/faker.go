package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	generator "github.com/Befous/DummyGenerator"
	"github.com/FreightTrackr/backend/helpers"
	"github.com/FreightTrackr/backend/models"
	"github.com/bxcodec/faker/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataAlamat struct {
	Kode_Pos       int    `json:"kode_pos,omitempty" bson:"kode_pos,omitempty"`
	Desa_Kelurahan string `json:"desa_kelurahan,omitempty" bson:"desa_kelurahan,omitempty"`
	Kecamatan      string `json:"kecamatan,omitempty" bson:"kecamatan,omitempty"`
	Kota_Kabupaten string `json:"kota_kabupaten,omitempty" bson:"kota_kabupaten,omitempty"`
	Ibu_Kota       string `json:"ibu_kota,omitempty" bson:"ibu_kota,omitempty"`
	Kode_Kota      string `json:"kode_kota,omitempty" bson:"kode_kota,omitempty"`
	Kode_Bandara   string `json:"kode_bandara,omitempty" bson:"kode_bandara,omitempty"`
	Provinsi       string `json:"provinsi,omitempty" bson:"provinsi,omitempty"`
	Alamat_Lengkap string `json:"alamat_lengkap,omitempty" bson:"alamat_lengkap,omitempty"`
}

func DummyUserGenerator(n int, mconn *mongo.Database) (string, error) {

	for i := 0; i < n; i++ {
		var user models.Users
		// Generate random user data with faker
		user.Username = faker.Username()
		user.Password = faker.Password()
		user.Nama = faker.Name()
		user.No_Telp = faker.Phonenumber()
		user.Email = faker.Email()
		user.Role = "user" // or any other roles such as "admin", "guest"

		// Ensure valid data according to the rules
		// Make sure username doesn't contain spaces and isn't too long
		if strings.Contains(user.Username, " ") || len(user.Username) > 20 {
			user.Username = faker.Username()
		}

		// Ensure password meets length constraints
		if len(user.Password) < 8 || len(user.Password) > 20 {
			user.Password = faker.Password()
		}

		// Ensure name length is within limits
		if len(user.Nama) > 55 {
			user.Nama = faker.Name()
		}

		// Generate a hash for the password using the same method as in Register
		hash, err := helpers.HashPassword(user.Password)
		if err != nil {
			return "Gagal hash", err
		}
		user.Password = hash

		InsertUser(mconn, "users", user)
	}

	return "Bergasil generate " + strconv.Itoa(n) + " data", nil
}

func DummyTransaksiGenerator(n int, mconn *mongo.Database) (string, error) {

	for i := 0; i < n; i++ {
		var transaksi models.Transaksi
		alamatpengirim := generator.GenerateRandomAlamat()
		alamatpenerima := generator.GenerateRandomAlamat()

		transaksi.No_Resi = "P" + strconv.Itoa(rand.Intn(100000000000))
		transaksi.Layanan = randomLayanan()
		transaksi.Isi_Kiriman = faker.Word()
		transaksi.Nama_Pengirim = faker.Name()
		transaksi.Alamat_Pengirim = alamatpengirim.Alamat_Lengkap
		transaksi.Kode_Pos_Pengirim = alamatpengirim.Kode_Pos
		transaksi.Kota_Asal = alamatpengirim.Kota_Kabupaten
		transaksi.Nama_Penerima = faker.Name()
		transaksi.Alamat_Penerima = alamatpenerima.Alamat_Lengkap
		transaksi.Kode_Pos_Penerima = alamatpenerima.Kode_Pos
		transaksi.Kota_Tujuan = alamatpenerima.Kota_Kabupaten
		transaksi.Berat_Kiriman = rand.Float64() * 10 // Random weight between 0 and 10 kg
		transaksi.Volumetrik = rand.Float64() * 500   // Random volumetrik between 0 and 500
		transaksi.Nilai_Barang = rand.Intn(10000000)  // Random value between 0 and 10 million
		transaksi.Biaya_Dasar = rand.Intn(100000)     // Random base cost between 0 and 100 thousand
		transaksi.Biaya_Pajak = int(float64(transaksi.Biaya_Dasar) * 0.11)
		transaksi.Biaya_Asuransi = int(float64(transaksi.Nilai_Barang) * 0.005)
		transaksi.Total_Biaya = transaksi.Biaya_Dasar + transaksi.Biaya_Pajak + transaksi.Biaya_Asuransi
		transaksi.Tanggal_Kirim = primitive.NewDateTimeFromTime(time.Now())
		transaksi.Tanggal_Terima = primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, 1))
		transaksi.Tanggal_Tenggat = primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, 1))
		transaksi.Status = randomStatus()
		transaksi.Tipe_Cod = randomTipeCod()
		transaksi.Status_Cod = randomStatusCod()
		transaksi.Sla = 1
		transaksi.Aktual_Sla = 1
		transaksi.Status_Sla = true
		transaksi.No_Pend_Kirim = fakeDigit()
		transaksi.No_Pend_Terima = fakeDigit()
		transaksi.Kode_Pelanggan = faker.Username()
		transaksi.Created_By.Username = faker.Username()
		transaksi.ID_History = faker.UUIDDigit()

		InsertTransaksi(mconn, "transaksi", transaksi)
	}

	return "Bergasil generate " + strconv.Itoa(n) + " data", nil
}

func DummyKantorGenerator(n int, mconn *mongo.Database) (string, error) {
	regionOptions := []int{1, 2, 3, 4, 5, 6}
	officeTypes := []string{"kcu", "kc", "kcp"}

	for i := 0; i < n; i++ {
		var kantor models.Kantor
		alamat := generator.GenerateRandomAlamat()
		kantor.No_Pend = fakeDigit()
		kantor.Tipe_Kantor = officeTypes[rand.Intn(len(officeTypes))]
		kantor.Nama_Kantor = "asdasd"
		kantor.Region_Kantor = regionOptions[rand.Intn(len(regionOptions))]
		kantor.Kota_Kantor = alamat.Kota_Kabupaten
		kantor.Kode_Pos_Kantor = alamat.Kode_Pos
		kantor.Alamat_Kantor = alamat.Alamat_Lengkap

		if kantor.Tipe_Kantor != "kcu" {
			kantor.No_Pend_Kcu = fakeDigit() // Isi dengan No_Pend jika Tipe_Kantor adalah kcu
		}

		if kantor.Tipe_Kantor != "kc" && kantor.Tipe_Kantor != "kcu" {
			kantor.No_Pend_Kc = fakeDigit()
		}

		InsertKantor(mconn, "kantor", kantor)
	}

	return "Bergasil generate " + strconv.Itoa(n) + " data", nil
}

func DummyPelangganGenerator(n int, mconn *mongo.Database) (string, error) {
	tipePelangganOptions := []string{"Retail", "Corporate", "Marketplace"}
	marketplaceNames := []string{"Shopee", "Tokopedia", "Bukalapak", "Lazada"}
	corporateNames := []string{"Mandiri", "BRI", "BCA", "CitiBank"}

	for i := 0; i < n; i++ {
		var pelanggan models.Pelanggan
		// Generate random data for Pelanggan
		pelanggan.Kode_Pelanggan = fakeDigit() // Assuming utils.RandomString generates a random string
		pelanggan.Tipe_Pelanggan = tipePelangganOptions[rand.Intn(len(tipePelangganOptions))]

		// If Tipe_Pelanggan is Retail, Nama_Pelanggan is empty or matches sender's name
		if pelanggan.Tipe_Pelanggan == "Retail" {
		} else if pelanggan.Tipe_Pelanggan == "Marketplace" {
			pelanggan.Nama_Pelanggan = marketplaceNames[rand.Intn(len(marketplaceNames))]
		} else if pelanggan.Tipe_Pelanggan == "Corporate" {
			pelanggan.Nama_Pelanggan = corporateNames[rand.Intn(len(corporateNames))]
		}

		// Insert the generated Pelanggan data into the database
		InsertPelanggan(mconn, "pelanggan", pelanggan)
	}

	return "Bergasil generate " + strconv.Itoa(n) + " data", nil
}

func DummyHistoryGenerator(n int, mconn *mongo.Database) (string, error) {
	usernames := []string{"user1", "user2", "user3", "user4"} // Example usernames

	for i := 0; i < n; i++ {
		var history models.History
		history.ID_History = fakeDigit() // Assuming utils.RandomString generates a random string

		for j := 0; j < rand.Intn(5)+1; j++ { // Generate between 1 and 5 status updates
			var lokasi models.Lokasi
			lokasi.Status = randomStatus()
			lokasi.Timestamp = time.Now()
			lokasi.Coordinate = []float64{rand.Float64() * 180.0, rand.Float64() * 180.0}
			lokasi.Catatan = "Paket masuk gudang"
			lokasi.Username = usernames[rand.Intn(len(usernames))]

			history.Lokasi = append(history.Lokasi, lokasi)
		}

		InsertHistory(mconn, "history", history)
	}

	return "Berhasil generate " + strconv.Itoa(n) + " data", nil
}

func randomLayanan() string {
	layanan := []string{"Reguler", "Cepat", "Express"}
	return layanan[rand.Intn(len(layanan))]
}

func randomStatus() string {
	status := []string{"delivered", "canceled", "returned", "inWarehouse", "inVehicle", "failed"}
	return status[rand.Intn(len(status))]
}

func randomTipeCod() string {
	tipeCod := []string{"no cod", "cod"}
	return tipeCod[rand.Intn(len(tipeCod))]
}

func randomStatusCod() string {
	statusCod := []string{"paid", "unpaid", "onProcess"}
	return statusCod[rand.Intn(len(statusCod))]
}

// Fungsi untuk menghasilkan kode pos palsu
func fakeZipCode() int {
	return rand.Intn(99999-10000) + 10000
}

// Fungsi untuk menghasilkan digit palsu
func fakeDigit() string {
	return strconv.Itoa(rand.Intn(1000000000))
}
