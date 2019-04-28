package main

import (
	"ELRA/database"
	"ELRA/globals"
	"ELRA/server"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Booting E.L.R.A. v" + globals.Version)
	globals.SetupGlobals()
	database.SetupDatabase()
	server.StartServer()
}
