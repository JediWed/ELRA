package database

import (
	"ELRA/globals"
	"ELRA/tools"
	"database/sql"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// DatabaseType describes the Database which is used
const DatabaseType = "sqlite3"

// SetupDatabase Creates the database if there is no one
func SetupDatabase() {
	if _, err := os.Stat(globals.Config.Database); os.IsNotExist(err) {
		log.Println("No database found. Creating...")
		db, err := sql.Open(DatabaseType, globals.Config.Database)
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
		insertuser, err := db.Prepare("INSERT INTO user (id, name, password, firstname, lastname, email, role) VALUES (?, ?, ?, ?, ?, ?, ?)")
		tools.CheckError(err)
		insertuser.Exec(1, "admin", string(passwordHash), "Admin", "Almighty", "admin@allnetworks.com", globals.RoleAdmin.ID)

		// Create Privileges
		createPrivileges, err := db.Prepare("CREATE TABLE privileges (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, description VARCHAR(255) NOT NULL);")
		tools.CheckError(err)
		createPrivileges.Exec()
		insertPrivilege, err := db.Prepare("INSERT INTO privileges (id, description) VALUES (?, ?)")
		tools.CheckError(err)
		for _, privilege := range globals.Privileges {
			insertPrivilege.Exec(privilege.ID, privilege.Description)
		}

		// Create UserPrivileges
		createUserPrivileges, err := db.Prepare("CREATE TABLE user_privileges (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, user INTEGER NOT NULL, privilege INTEGER NOT NULL, FOREIGN KEY (user) REFERENCES user(id) ON DELETE CASCADE, FOREIGN KEY (privilege) REFERENCES privileges(id) ON DELETE RESTRICT)")
		tools.CheckError(err)
		createUserPrivileges.Exec()
		insertUserPrivilege, err := db.Prepare("INSERT INTO user_privileges (user, privilege) VALUES (?, ?)")
		tools.CheckError(err)
		for _, privilege := range globals.Privileges {
			insertUserPrivilege.Exec(1, privilege.ID)
		}
	} else {
		log.Println("Loading database...")
	}
}
