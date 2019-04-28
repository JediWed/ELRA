package server

import (
	"ELRA/authorization"
	"ELRA/database"
	"ELRA/structs"
	"encoding/json"
	"log"
	"net/http"
)

func Login(response http.ResponseWriter, request *http.Request) {
	SetupCORS(&response, request)
	if (*request).Method == "OPTIONS" {
		return
	}

	var loginRequest structs.LoginRequest
	err := json.NewDecoder(request.Body).Decode(&loginRequest)
	if err != nil {
		log.Print("Login attempt with insufficient parameters.")
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	json, err := database.Login(loginRequest)

	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(json)
}

func UpdatePassword(response http.ResponseWriter, request *http.Request) {
	SetupCORS(&response, request)
	userid, _, err := authorization.ParseUserIDAndRole(request)

	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	var updateRequest structs.UpdatePasswordRequest

	err = json.NewDecoder(request.Body).Decode(&updateRequest)
	if err != nil {
		log.Print("Change Password attempt with insufficient parameters. UserID ", userid)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	updateRequest.UserID = userid

	err = database.UpdatePassword(updateRequest)

	if err != nil {
		log.Print("Could not change password for UserID ", userid)
		response.WriteHeader(http.StatusUnauthorized)
	}

	log.Print("Password of UserID ", userid, " was changed")
}

func UpdateUsername(response http.ResponseWriter, request *http.Request) {
	SetupCORS(&response, request)
	userid, _, err := authorization.ParseUserIDAndRole(request)

	if err != nil {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	var updateRequest structs.UpdateUsernameRequest

	err = json.NewDecoder(request.Body).Decode(&updateRequest)
	if err != nil {
		log.Print("Change Username attempt with insufficient parameters. UserID ", userid)
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	updateRequest.UserID = userid

	err = database.UpdateUsername(updateRequest)

	if err != nil {
		log.Print("Could not change username for UserID ", userid)
		response.WriteHeader(http.StatusUnauthorized)
	}

	log.Print("Username of UserID ", userid, " was changed to ", updateRequest.Username)

}
