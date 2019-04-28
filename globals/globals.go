package globals

import (
	"ELRA/structs"
	"ELRA/tools"
	"encoding/json"
	"log"
	"os"
)

const Version = "1.0"

// Roles
var RoleAdmin = structs.DescriptionType{ID: 1, Description: "Admin"}
var RoleUser = structs.DescriptionType{ID: 2, Description: "User"}

var Config structs.Configuration

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

}