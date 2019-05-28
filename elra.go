package main

import (
	"ELRA/database"
	"ELRA/globals"
	"ELRA/lnd"
	"ELRA/server"
	"ELRA/tools"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Booting E.L.R.A. v" + globals.Version)
	globals.SetupGlobals()
	database.SetupDatabase()
	lnd.SetupLND()
	go tools.StartPriceDaemon(globals.Config)
	server.StartServer()
}
