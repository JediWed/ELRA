package server

import (
	"ELRA/authorization"
	"ELRA/database"
	"ELRA/structs"
	"encoding/json"
	"log"
	"net/http"
)

// LoginEndpoint is the Endpoint for Login
const LoginEndpoint = "/user/login"

// UpdatePasswordEndpoint is the Endpoint for UpdatePassword
const UpdatePasswordEndpoint = "/user/updatePassword"

// UpdateUsernameEndpoint is the Endpoint for UpdateUsername
const UpdateUsernameEndpoint = "/user/updateUsername"

// Login is the server sided function to handle login endpoint
func Login(response http.ResponseWriter, request *http.Request) {
	database.AccessLog(request.Header.Get("X-Forwarded-For"), LoginEndpoint)
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

// UpdatePassword is the server sided function to handle update password endpoint
func UpdatePassword(response http.ResponseWriter, request *http.Request) {
	database.AccessLog(request.Header.Get("X-Forwarded-For"), UpdatePasswordEndpoint)
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

// UpdateUsername is the server sided function to handle update username endpoint
func UpdateUsername(response http.ResponseWriter, request *http.Request) {
	database.AccessLog(request.Header.Get("X-Forwarded-For"), UpdateUsernameEndpoint)
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
