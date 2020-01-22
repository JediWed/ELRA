package server

import (
	"ELRA/authorization"
	"ELRA/globals"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupCORS sets the CORS with Access-Control-Allow-Origin = *
func SetupCORS(w *http.ResponseWriter, _ *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// StartServer starts the whole API Server
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

	// root
	router.HandleFunc(GetVersionEndpoint, GetVersion).Methods("GET")

	// /invoice
	router.HandleFunc(GetInvoiceEndpoint, GetInvoice).Methods("GET")
	router.HandleFunc(CreateInvoiceEndpoint, CreateInvoice).Methods("GET")

	// /user
	router.HandleFunc(LoginEndpoint, Login).Methods("POST")
	router.Handle(UpdatePasswordEndpoint, authorization.CheckAuthorization(UpdatePassword)).Methods("POST")
	router.Handle(UpdateUsernameEndpoint, authorization.CheckAuthorization(UpdateUsername)).Methods("POST")

	log.Print("Starting REST API on: ", globals.Config.AllowedHost, ":", globals.Config.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", globals.Config.AllowedHost, globals.Config.Port), router)
}
