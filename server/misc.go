package server

import (
	"ELRA/database"
	"ELRA/globals"
	"ELRA/tools"
	"fmt"
	"net/http"
)

// GetVersionEndpoint is the Endpoint for GetVersion
const GetVersionEndpoint = "/version"

// GetVersion returns Version of ELRA
func GetVersion(response http.ResponseWriter, request *http.Request) {
	database.AccessLog(tools.ExtractIPAddressFromRemoteAddr(request.RemoteAddr), GetVersionEndpoint)
	SetupCORS(&response, request)
	response.Write([]byte(fmt.Sprintf("{\"version\": \"%s\"}", globals.Version)))
}
