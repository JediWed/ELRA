package globals

import (
	"ELRA/structs"
	"ELRA/tools"
	"encoding/json"
	"os"
)

const Version = "1.0"

var Config structs.Configuration

func SetupGlobals() {
	configFile, err := os.Open("./config.json")
	tools.CheckError(err)
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&Config)
	tools.CheckError(err)

	if Config.Port <= 0 || Config.Port > 65535 {
		Config.Port = 8181
	}

	if Config.Pepper == "" {
		Config.Pepper = "C6yCmdDaCDkxiIrRBngJ2mhil9ihHnM6rDP6Pp7Zn4"
	}

	if Config.SigningKey == "" {
		Config.Pepper = "lss8hzIsJbXCP4g33yp12LtsrMJCehTK"
	}

	if Config.Database == "" {
		Config.Database = "./elra.db"
	}

}
