package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserStruct struct {
	Username string
	Forename string
	Lastname string
	Email    string
}

func api_versions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")
	w.Write([]byte("/api/v0,/api/v1"))
}

func api_v0_version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v0")
	w.Write([]byte("v0.1234"))
}

func api_v1_version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "v1")
	w.Write([]byte("v1.2345"))
}

func api_user_profile(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//name := params["name"]
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Api-Service", "-")
	w.Write([]byte("Hello World"))
}

// curl -X POST -d '{"username":"dieter", "forename":"Dieter", "lastname": "Reuter", "email":"dieter.reuter@me.com"}' http://127.0.0.1:10777/api/v1/users
func api_v1_users(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println("Body=", string(body))
	var user UserStruct
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	d, _ := json.Marshal(user)
	log.Println("JSON=", string(d))
}
