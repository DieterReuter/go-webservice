package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func init() {
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
// curl -X POST -d '{"username":"hugo", "forename":"Hugo", "lastname": "Boss", "email":"hugo.boss@hugoboss.com"}' http://127.0.0.1:10777/api/v1/users
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
