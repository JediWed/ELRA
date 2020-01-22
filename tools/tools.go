package tools

import (
	"log"
	"net/http"
	"strings"
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
	}

	ipAddress = request.Header.Get("X-Forwarded-For")

	if ipAddress != "" {
		return ExtractIPAddressFromRemoteAddr(ipAddress)
	}

	ipAddress = request.RemoteAddr

	if ipAddress != "" {
		return ExtractIPAddressFromRemoteAddr(ipAddress)
	}

	return "0.0.0.0"
}

// ExtractIPAddressFromRemoteAddr extracts IP Address from a Remote Address
func ExtractIPAddressFromRemoteAddr(remoteAddr string) string {
	ipIndex := strings.LastIndex(remoteAddr, ":")
	if ipIndex != -1 {
		ip := remoteAddr[0:ipIndex]
		return ip
	}
	return remoteAddr
}
