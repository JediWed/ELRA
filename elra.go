package main

import (
	"ELRA/globals"
	"log"
)

func main() {
	log.Println("Booting E.L.R.A. v" + globals.Version)
	globals.SetupGlobals()
}
