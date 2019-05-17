package database

import (
	"ELRA/authorization"
	"ELRA/globals"
	"ELRA/structs"
	"ELRA/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Login(loginRequest structs.LoginRequest) ([]byte, error) {
	db, err := sql.Open(DatabaseType, globals.Config.Database)
	tools.CheckError(err)
	defer db.Close()

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM user WHERE name = ?", loginRequest.Username)
	err = row.Scan(&count)

	if count != 1 {
		log.Print("Unsuccessful Login attempt with username ", loginRequest.Username)
		return nil, fmt.Errorf("Unsuccessful Login attempt with username " + loginRequest.Username)
	}
	tools.CheckError(err)

	row = db.QueryRow("SELECT * FROM user WHERE name = ?", loginRequest.Username)
	var id int
	var name string
	var password string
	var firstname string
	var lastname string
	var email string
	var role int

	row.Scan(&id, &name, &password, &firstname, &lastname, &email, &role)

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(loginRequest.Password+globals.Config.Pepper))

	if err != nil {
		log.Print("Unsuccessful Login attempt with username (wrong password) ", loginRequest.Username)
		log.Print(err)
		return nil, fmt.Errorf("Unsuccessful Login attempt with username (wrong password) " + loginRequest.Username)
	}

	token, err := authorization.CreateJWTToken(id, role, globals.Config)

	if err != nil {
		log.Print("Could not create JWT Token.")
		log.Print(err)
		return nil, fmt.Errorf("Could not create JWT Token.")
	}

	loginResponse := structs.LoginResponse{ID: id, Name: name, Firstname: firstname, Lastname: lastname, Email: email, Role: role, Token: token}

	json, err := json.Marshal(loginResponse)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Print("Successful login of user " + loginRequest.Username)

	return json, nil
}

func UpdatePassword(updateRequest structs.UpdatePasswordRequest) error {

	db, err := sql.Open(DatabaseType, globals.Config.Database)
	tools.CheckError(err)
	defer db.Close()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(updateRequest.Password+globals.Config.Pepper), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	updatePassword, err := db.Prepare("UPDATE user SET password = ? WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = updatePassword.Exec(string(passwordHash), updateRequest.UserID)

	return err
}

func UpdateUsername(updateRequest structs.UpdateUsernameRequest) error {

	db, err := sql.Open(DatabaseType, globals.Config.Database)
	tools.CheckError(err)
	defer db.Close()

	if err != nil {
		return err
	}

	updatePassword, err := db.Prepare("UPDATE user SET name = ? WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = updatePassword.Exec(updateRequest.Username, updateRequest.UserID)

	return err
}
