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

func generateRandomDate(start time.Time, end time.Time) time.Time {
	diff := end.Sub(start)
	randDuration := time.Duration(rand.Int63n(int64(diff)))
	return start.Add(randDuration)
}

func DummyTransaksiGenerator(n int, mconn *mongo.Database) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	layanan := []string{"Nextday", "Reguler"}
	isi_kiriman := []string{"Dokumen", "Paket"}
	tipe_cod := []string{"cod", "nonCod"}
	status_cod := []string{"sudah_setor", "belum_setor"}
	status := []string{"delivered", "canceled", "returned", "inWarehouse", "inVehicle", "failed", "paid"}
	startDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Now()
	for i := 0; i < n; i++ {
		var transaksi models.Transaksi

		alamatpengirim := generator.GenerateRandomAlamat()
		alamatpenerima := generator.GenerateRandomAlamat()

		data_status := status[r.Intn(len(status))]
		data_tipe_cod := tipe_cod[r.Intn(len(tipe_cod))]
		tanggal_kirim := generateRandomDate(startDate, endDate)
		transaksi.Tanggal_Kirim = primitive.NewDateTimeFromTime(tanggal_kirim)
		transaksi.Sla = rand.Intn(4) + 2

		if data_status == "delivered" {
			var tanggal_antaran_pertama, tanggal_terima time.Time
			tanggal_antaran_pertama = tanggal_kirim.Add(time.Duration(rand.Intn(4)+3) * 24 * time.Hour)
			tanggal_terima = tanggal_antaran_pertama.Add(time.Duration(rand.Intn(3)) * 24 * time.Hour)
			transaksi.Tanggal_Antaran_Pertama = primitive.NewDateTimeFromTime(tanggal_antaran_pertama)
			transaksi.Tanggal_Terima = primitive.NewDateTimeFromTime(tanggal_terima)
			transaksi.Aktual_Sla = int(tanggal_antaran_pertama.Sub(tanggal_kirim).Hours() / 24)
			transaksi.Status_Sla = transaksi.Aktual_Sla <= transaksi.Sla
		}
		if data_tipe_cod == "cod" {
			if data_status == "delivered" {
				transaksi.Status_Cod = status_cod[r.Intn(len(status_cod))]
			} else {
				transaksi.Status_Cod = "belum_setor"
			}
		}

		year := tanggal_kirim.Year()
		data_layanan := layanan[r.Intn(len(layanan))]
		data_isi_kiriman := isi_kiriman[r.Intn(len(isi_kiriman))]
		data_berat_kiriman := float64(rand.Intn(100)) + (float64(rand.Intn(9)+1) / 10)

		data_volumetrik := float64(rand.Intn(491)+10) + float64(rand.Intn(10))/10.0

		base := (rand.Intn(90) + 10) * 1000
		extra := []int{0, 500}[rand.Intn(2)]
		data_biaya_dasar := base + extra

		transaksi.No_Resi = "P" + strconv.Itoa(year)[2:] + strconv.Itoa(rand.Intn(10000000))
		transaksi.Layanan = data_layanan
		transaksi.Isi_Kiriman = data_isi_kiriman
		transaksi.Nama_Pengirim = faker.Name()
		transaksi.Alamat_Pengirim = alamatpengirim.Alamat_Lengkap
		transaksi.Kode_Pos_Pengirim = alamatpengirim.Kode_Pos
		transaksi.Kota_Asal = alamatpengirim.Kota_Kabupaten
		transaksi.Nama_Penerima = faker.Name()
		transaksi.Alamat_Penerima = alamatpenerima.Alamat_Lengkap
		transaksi.Kode_Pos_Penerima = alamatpenerima.Kode_Pos
		transaksi.Kota_Tujuan = alamatpenerima.Kota_Kabupaten
		transaksi.Berat_Kiriman = data_berat_kiriman
		transaksi.Volumetrik = data_volumetrik
		transaksi.Biaya_Dasar = data_biaya_dasar
		transaksi.Biaya_Pajak = int(float64(data_biaya_dasar) * 0.11)

		if rand.Intn(2) == 0 {
			transaksi.Nilai_Barang = 0
			transaksi.Biaya_Asuransi = 0
		} else {
			transaksi.Nilai_Barang = (100 + rand.Intn(901)) * 1000
			transaksi.Biaya_Asuransi = int(float64(transaksi.Nilai_Barang) * 0.005)
		}

		transaksi.Total_Biaya = transaksi.Biaya_Dasar + transaksi.Biaya_Pajak + transaksi.Biaya_Asuransi
		transaksi.Status = data_status
		transaksi.Tipe_Cod = data_tipe_cod
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
	tipePelangganOptions := []string{"Corporate", "Marketplace"}
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

func randomStatus() string {
	status := []string{"delivered", "canceled", "returned", "inWarehouse", "inVehicle", "failed"}
	return status[rand.Intn(len(status))]
}

// Fungsi untuk menghasilkan digit palsu
func fakeDigit() string {
	return strconv.Itoa(rand.Intn(1000000000))
}
