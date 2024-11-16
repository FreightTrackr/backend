package testing

import (
	"fmt"
	"testing"

	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/utils"
)

func TestRegisterDummyUsers(t *testing.T) {
	config.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyUserGenerator(100, mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}

func TestGenerateDataDummyTransaksi(t *testing.T) {
	config.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyTransaksiGenerator(50, mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
func TestGenerateDataDummyKantor(t *testing.T) {
	config.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyKantorGenerator(50, mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
func TestGenerateDataDummyPelanggan(t *testing.T) {
	config.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyKantorGenerator(50, mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
func TestGenerateDataDummyHistory(t *testing.T) {
	config.LoadEnv("../.env")
	mconn := utils.SetConnection()
	apalah, err := utils.DummyHistoryGenerator(50, mconn)
	if err != nil {
		t.Fatalf("Error generating dummy users: %v", err)
	}
	fmt.Println(apalah)
}
