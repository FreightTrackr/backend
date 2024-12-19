package main

import (
	"log"

	"github.com/FreightTrackr/backend/middleware"
	"github.com/FreightTrackr/backend/utils"
)

func main() {
	middleware.LoadEnv(".env")

	mconn := utils.SetConnection()
	// _, err := utils.DummyUserGenerator(179, mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy users: %v", err)
	// }
	_, err := utils.DummyTransaksiGenerator(1000, mconn)
	if err != nil {
		log.Fatalf("Error generating dummy transaksi: %v", err)
	}
	// _, err = utils.DummyKantorGenerator(mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy kantor: %v", err)
	// }
	// _, err = utils.DummyPelangganGenerator(mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy pelanggan: %v", err)
	// }
	// _, err = utils.DummyHistoryGenerator(30, mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy history: %v", err)
	// }
}
