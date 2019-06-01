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

// AccessLimitExceeded checks a max access count for a specific ip and endpoint in 60 seconds
func AccessLimitExceeded(ip string, endpoint string, maxPerMinute int, maxPerHour int) (bool, error) {
	db, err := sql.Open(DatabaseType, globals.Config.Database)
	tools.CheckError(err)
	defer db.Close()

	var countMinute int
	row := db.QueryRow("SELECT count(*) FROM access_log WHERE date > datetime('now', '-60 seconds') AND ip = ? AND endpoint = ?", ip, endpoint)
	err = row.Scan(&countMinute)

	if err != nil {
		return true, err
	}

	var countHour int
	row = db.QueryRow("SELECT count(*) FROM access_log WHERE date > datetime('now', '-60 minutes') AND ip = ? AND endpoint = ?", ip, endpoint)
	err = row.Scan(&countHour)

	if err != nil {
		return true, err
	}

	return (countMinute >= maxPerMinute || countHour >= maxPerHour), nil
}
