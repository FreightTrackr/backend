package main

import (
	"fmt"
	"net/http"

	"github.com/FreightTrackr/backend/config"
	"github.com/FreightTrackr/backend/routes"
)

func main() {
	config.LoadEnv(".env")
	app := http.NewServeMux()
	routes.StdRoute(app)
	server := http.Server{
		Addr:    ":3000",
		Handler: app,
	}
	fmt.Println("Starting server on port :3000")
	server.ListenAndServe()
}
