package server

import (
	"ELRA/globals"
	"fmt"
	"net/http"
)

// GetVersion returns Version of ELRA
func GetVersion(response http.ResponseWriter, request *http.Request) {
	SetupCORS(&response, request)
	response.Write([]byte(fmt.Sprintf("{\"version\": \"%s\"}", globals.Version)))
}
