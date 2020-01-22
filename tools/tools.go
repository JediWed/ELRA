package tools

import (
	"log"
	"strings"
	"net/http"
)

// CheckError checks an Error and exits application with a fatal error
func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}

// ExtractIPAddressFromRequest extracts IP Address from a Request
func ExtractIPAddressFromRequest(request *http.Request) string {
	ipAddress := request.Header.Get("X-Real-IP")

	if ipAddress != "" {
		return ExtractIPAddressFromRemoteAddr(ipAddress)
	} else {
		ipAddress = request.Header.Get("X-Forwarded-For")
	}

	if ipAddress != "" {
		return ipAddress
	} else {
		ipAddress = request.RemoteAddr
	}

	if ipAddress != "" {
		return ExtractIPAddressFromRemoteAddr(ipAddress)
	}

	return "0.0.0.0"
}

// ExtractIPAddressFromRemoteAddr extracts IP Address from a Remote Address
func ExtractIPAddressFromRemoteAddr(remoteAddr string) string {
	log.Print(remoteAddr)
	ipIndex := strings.LastIndex(remoteAddr, ":")
	ip := remoteAddr[0:ipIndex]
	return ip
}
