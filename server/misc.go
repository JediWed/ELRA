package server

import (
	"ELRA/database"
	"ELRA/globals"
	"fmt"
	"net/http"
)

// GetVersionEndpoint is the Endpoint for GetVersion
const GetVersionEndpoint = "/version"

// GetVersion returns Version of ELRA
func GetVersion(response http.ResponseWriter, request *http.Request) {
	database.AccessLog(request.Header.Get("X-Forwarded-For"), GetVersionEndpoint)
	SetupCORS(&response, request)
	response.Write([]byte(fmt.Sprintf("{\"version\": \"%s\"}", globals.Version)))
}
