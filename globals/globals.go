package globals

import (
	"ELRA/structs"
	"ELRA/tools"
	"encoding/json"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

// Version is the global version of ELRA
const Version = "Alpha 01"

// Roles

// RoleAdmin is the role for Admins
var RoleAdmin = structs.DescriptionType{ID: 1, Description: "Admin"}

// RoleUser ist the role für Users
var RoleUser = structs.DescriptionType{ID: 2, Description: "User"}

// Privilege is a struct for privileges
type Privilege structs.DescriptionType

var (
	// GetInfoPrivilege is the privilege to use GetInfo
	GetInfoPrivilege = Privilege{ID: 1, Description: "GetInfo"}
)

// Privileges is an array of all available privileges
var Privileges = []Privilege{GetInfoPrivilege}

// Config is the global configuration singleton of ELRA
var Config structs.Configuration

// SetupGlobals sets up all global configurations
func SetupGlobals() {
	configFile, err := os.Open("./config.json")
	tools.CheckError(err)
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&Config)
	tools.CheckError(err)

	if Config.Port <= 0 || Config.Port > 65535 {
		log.Println("No configured Port was found. Setting to 8181")
		Config.Port = 8181
	}

	if Config.Pepper == "" {
		log.Println("No configured Pepper was found. Creating Default Pepper. This is very dangerous. You should set a Signing Key in your config.")
		Config.Pepper = "C6yCmdDaCDkxiIrRBngJ2mhil9ihHnM6rDP6Pp7Zn4"
	}

	if Config.SigningKey == "" {
		log.Println("No configured Signing Kex was found. Creating Default Signing Key. This is very dangerous. You should set a Signing Key in your config.")
		Config.SigningKey = "lss8hzIsJbXCP4g33yp12LtsrMJCehTK"
	}

	if Config.Database == "" {
		log.Println("No Database was configured. Setting to default database elra.db.")
		Config.Database = "./elra.db"
	}

	if Config.Macaroon == "" {
		log.Println("No Macaroon was configured. Setting to local admin.macaroon.")
		Config.Macaroon = "./admin.macaroon"
	}

	if Config.TLS == "" {
		log.Println("No TLS was configured. Setting to local tls.cert.")
		Config.Macaroon = "./tls.cert"
	}

	if Config.LightningServer == "" {
		log.Println("No Lightning Server was configured. Setting to 127.0.0.1.")
		Config.LightningServer = "127.0.0.1"
	} else {
		Config.LightningServer = strings.ToLower(Config.LightningServer)
		ipRegEx, _ := regexp.Compile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
		domainRegEx, _ := regexp.Compile(`^(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`)
		if domainRegEx.MatchString(Config.LightningServer) || Config.LightningServer == "localhost" {
			ips, err := net.LookupIP(Config.LightningServer)
			tools.CheckError(err)
			Config.LightningServer = ips[0].String()
		} else if !ipRegEx.MatchString(Config.LightningServer) && Config.LightningServer != "localhost" {
			log.Fatal("Your Lightning Server URL (" + Config.LightningServer + ") is wrong. Please use FQDN Format!")
		}
	}

	if Config.LightninggRPCPort <= 0 || Config.LightninggRPCPort > 65535 {
		log.Println("No configured Lighting Port was found. Setting to 10009")
		Config.Port = 10009
	}

	if Config.BitcoinPriceAPI == "" {
		log.Println("No Bitcoin Price API was found. Setting to https://api.bugs-ev.de/v1/bitcoinPrice")
		Config.BitcoinPriceAPI = "https://api.bugs-ev.de/v1/bitcoinPrice"
		Config.BitcoinPriceAPIKeyword = "bitcoin"
	} else if Config.BitcoinPriceAPIKeyword == "" {
		log.Fatal("No Bitcoin Price API Keyword was found. Exiting...")
	}
}
