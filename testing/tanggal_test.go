package testing

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestTanggal(t *testing.T) {
	tanggal_kirim := time.Now()
	var tanggal_antaran_pertama, tanggal_terima time.Time

	tanggal_antaran_pertama = tanggal_kirim.Add(time.Duration(rand.Intn(4)+3) * 24 * time.Hour)
	tanggal_terima = tanggal_antaran_pertama.Add(time.Duration(rand.Intn(3)) * 24 * time.Hour)

	fmt.Println("Tanggal Kirim:", tanggal_kirim)
	fmt.Println("Tanggal Antaran Pertama:", tanggal_antaran_pertama)
	fmt.Println("Tanggal Terima:", tanggal_terima)
}
