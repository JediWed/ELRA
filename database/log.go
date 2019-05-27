package database

import (
	"ELRA/globals"
	"ELRA/tools"
	"database/sql"
)

// AccessLog creates an access log entry for an endpoint with the ip address of the requester
func AccessLog(ip string, endpoint string) error {
	db, err := sql.Open(DatabaseType, globals.Config.Database)
	tools.CheckError(err)
	defer db.Close()

	insertLog, err := db.Prepare("INSERT INTO access_log (ip, endpoint) VALUES (?, ?)")

	if err != nil {
		return err
	}

	_, err = insertLog.Exec(ip, endpoint)

	return err
}
