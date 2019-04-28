package database

import (
	"ELRA/globals"
	"ELRA/tools"
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// SetupDatabase Creates the database if there is no one
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

		// Create User
		createUser, err := db.Prepare("CREATE TABLE user (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, name VARCHAR(255) NOT NULL, password VARCHAR(255) NOT NULL, firstname VARCHAR(255) NOT NULL, lastname VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, role INTEGER NOT NULL, FOREIGN KEY (role) REFERENCES user_role(id));")
		tools.CheckError(err)
		createUser.Exec()
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("private"+globals.Config.Pepper), bcrypt.DefaultCost)
		insertuser, err := db.Prepare("INSERT INTO user (name, password, firstname, lastname, email, role) VALUES (?, ?, ?, ?, ?, ?)")
		tools.CheckError(err)
		insertuser.Exec("admin", string(passwordHash), "Admin", "Almighty", "admin@allnetworks.com", globals.RoleAdmin.ID)

	} else {
		log.Println("Loading database...")
	}
}
