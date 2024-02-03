package main

import (
	"github.com/abhilash26/gohyper/internal/router"
	"github.com/abhilash26/gohyper/internal/server"
	"github.com/abhilash26/gohyper/internal/storage"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Get Web Routes
	routes := router.SetupWebRoutes()
	storage.InitSQLite()
	// Start Web Server
	server.Start(routes)
}
