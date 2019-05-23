package server

import (
	"ELRA/authorization"
	"ELRA/globals"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupCORS(w *http.ResponseWriter, _ *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func StartServer() {
	log.Print("Preparing REST API")

	router := mux.NewRouter()

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	router.HandleFunc("/version", GetVersion).Methods("GET")
	router.HandleFunc("/invoice/createInvoice", CreateInvoice).Methods("GET")
	router.HandleFunc("/user/login", Login).Methods("POST")
	router.Handle("/user/updatePassword", authorization.CheckAuthorization(UpdatePassword)).Methods("POST")
	router.Handle("/user/updateUsername", authorization.CheckAuthorization(UpdateUsername)).Methods("POST")

	log.Print("Starting REST API on: ", globals.Config.AllowedHost, ":", globals.Config.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", globals.Config.AllowedHost, globals.Config.Port), router)
}
