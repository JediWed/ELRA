package tools

import (
	"log"
)

// CheckError checks an Error and exits application with a fatal error
func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}
