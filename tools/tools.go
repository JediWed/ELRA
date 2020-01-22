package tools

import (
	"log"
	"strings"
)

// CheckError checks an Error and exits application with a fatal error
func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}

// ExtractIPAddressFromRemoteAddr extracts IP Address from a Remote Address
func ExtractIPAddressFromRemoteAddr(remoteAddr string) string {
	ipIndex := strings.LastIndex(remoteAddr, ":")
	ip := remoteAddr[0:ipIndex]
	return ip
}
