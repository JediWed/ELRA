package tools

import (
	"log"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}

func ExtractIPAddressFromRemoteAddr(remoteAddr string) string {
	ipIndex := strings.LastIndex(remoteAddr, ":")
	ip := remoteAddr[0:ipIndex]
	return ip
}
