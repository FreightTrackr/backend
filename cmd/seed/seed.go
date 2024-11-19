package main

import (
	"log"

	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/utils"
)

func main() {
	config.LoadEnv(".env")

	mconn := utils.SetConnection()
	_, err := utils.DummyUserGenerator(179, mconn)
	if err != nil {
		log.Fatalf("Error generating dummy users: %v", err)
	}
	// _, err = utils.DummyTransaksiGenerator(50, mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy users: %v", err)
	// }
	// _, err = utils.DummyKantorGenerator(20, mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy users: %v", err)
	// }
	// _, err = utils.DummyPelangganGenerator(100, mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy users: %v", err)
	// }
	// _, err = utils.DummyHistoryGenerator(30, mconn)
	// if err != nil {
	// 	log.Fatalf("Error generating dummy users: %v", err)
	// }
}
