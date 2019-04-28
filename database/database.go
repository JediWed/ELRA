package database

import (
	"ELRA/globals"
	"ELRA/tools"
	"database/sql"
	"log"
	"os"
)

func SetupDatabase() {
	if _, err := os.Stat(globals.Config.Database); os.IsNotExist(err) {
		log.Println("No database found. Creating...")
		db, err := sql.Open("sqlite3", globals.Config.Database)
		tools.CheckError(err)
		defer db.Close()

		// Create User Roles
		createUserRole, err := db.Prepare("CREATE TABLE user_role (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, description VARCHAR(255) NOT NULL);")
		tools.CheckError(err)
		createUserRole.Exec()
		insertUserRoles, err := db.Prepare("INSERT INTO user_role (id, description) VALUES (?, ?)")
		tools.CheckError(err)
		insertUserRoles.Exec(globals.RoleAdmin.ID, globals.RoleAdmin.Description)
		insertUserRoles.Exec(globals.RoleUser.ID, globals.RoleUser.Description)

	} else {
		log.Println("Loading database...")
	}
}
