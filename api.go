package main

import (
	"net/http"
)

func api_versions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/api/v0,/api/v1"))
}

func api_v0_version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v0.1234"))
}

func api_v1_version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1.2345"))
}

func api_user_profile(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//name := params["name"]
	w.Write([]byte("Hello World"))
}
