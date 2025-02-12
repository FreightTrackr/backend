package testing

import (
	"fmt"
	"testing"

	"github.com/FreightTrackr/backend/middleware"
	"github.com/FreightTrackr/backend/utils"
)

func TestRegisterDummyUsers(t *testing.T) {
	middleware.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyUserGenerator(50, mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}

func TestGenerateDataDummyTransaksi(t *testing.T) {
	middleware.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyTransaksiGenerator(mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
func TestGenerateDataDummyKantor(t *testing.T) {
	middleware.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyKantorGenerator(mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
func TestGenerateDataDummyPelanggan(t *testing.T) {
	middleware.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyPelangganGenerator(mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
func TestGenerateDataDummyHistory(t *testing.T) {
	middleware.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyHistoryGenerator(mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
