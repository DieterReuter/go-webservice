package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var API_SERVERPORT string
var CONFIG_FILE string
var DATASTORE_DIR string

func main() {
	fmt.Println("API_SERVERPORT=", API_SERVERPORT)
	fmt.Println("CONFIG_FILE=", CONFIG_FILE)
	fmt.Println("DATASTORE_DIR=", DATASTORE_DIR)

	rtr := mux.NewRouter()
	rtr.HandleFunc("/versions", api_versions).Methods("GET")
	rtr.HandleFunc("/api", api_versions).Methods("GET")
	rtr.HandleFunc("/api/v0", api_v0_version).Methods("GET")
	rtr.HandleFunc("/api/v0/version", api_v0_version).Methods("GET")
	rtr.HandleFunc("/api/v1", api_v1_version).Methods("GET")
	rtr.HandleFunc("/api/v1/version", api_v1_version).Methods("GET")
	rtr.HandleFunc("/api/v1/users", api_v1_users).Methods("POST")

	rtr.HandleFunc("/user/dieter/profile", api_user_profile).Methods("GET")
	http.Handle("/", rtr)

	log.Println("Listening..." + API_SERVERPORT)
	http.ListenAndServe(API_SERVERPORT, nil)
}
