package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"github.com/golang_project_01_server/services"
)

type Client struct {
	Id             string `json:"id"`
	ClientName     string `json:"clientName"`
	ContractLength int    `json:"contractLength"`
}

type Project struct {
	Id          string `json:"id"`
	ProjectName string `json:"projectName"`
}

type SOW struct {
	Id       string `json:"id"`
	WorkName string `json:"workName"`
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

var newEmployee Employee
var newProject Project
var newSow SOW
var newClient Client
var token string

func main() {

	empContent, empErr := ioutil.ReadFile("employees.json")
	if empErr != nil {
		fmt.Println("could not read the file")
	}

	empErr2 := json.Unmarshal(empContent, &newEmployee)
	if empErr2 != nil {
		fmt.Println("unmarshal error")
	}

	projectContent, proErr := ioutil.ReadFile("projects.json")
	if proErr != nil {
		fmt.Println("could not read the file")
	}

	proErr2 := json.Unmarshal(projectContent, &newProject)
	if proErr2 != nil {
		fmt.Println("unmarshal error")
	}

	sowContent, sowErr := ioutil.ReadFile("sow.json")
	if sowErr != nil {
		fmt.Println("could not read the file")
	}

	sowErr2 := json.Unmarshal(sowContent, &newSow)
	if sowErr2 != nil {
		fmt.Println("unmarshal error")
	}

	clientContent, clientErr := ioutil.ReadFile("clients.json")
	if clientErr != nil {
		fmt.Println("could not read the file")
	}

	clientErr2 := json.Unmarshal(clientContent, &newClient)
	if clientErr2 != nil {
		fmt.Println("unmarshal error")
	}

	r := mux.NewRouter()

    r.HandleFunc("/auth", services.Auth).Methods("POST")
	r.HandleFunc("/login", BasicAuth(login)).Methods("GET")
	r.HandleFunc("/employees", ValidateTokenMiddleware(getEmployees)).Methods("GET")
	r.HandleFunc("/projects", ValidateTokenMiddleware(getProjects)).Methods("GET")
	r.HandleFunc("/sow", ValidateTokenMiddleware(getSow)).Methods("GET")
	r.HandleFunc("/clients", ValidateTokenMiddleware(getClients)).Methods("GET")

	http.ListenAndServe(":8081", r)

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmployee)
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProject)
}

func getSow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newSow)
}

func getClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newClient)
}

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
		fmt.Println("ok: ", ok)
		token = password
		if !ok || !checkUserAndPassword(username, password) {
			w.Header().Set("x-auth-token", "invalid")
			w.WriteHeader(401)
			w.Write([]byte("not authorized"))
			handler.ServeHTTP(w, r)
			return
		} else {

			handler.ServeHTTP(w, r)
		}

	})
}

func checkUserAndPassword(username, password string) bool {
	return username == "cc" && password == "password"
}

func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token == "password" {
			next(w, r)
		} else {
			w.Header().Set("x-auth-token", "invalid")
			w.WriteHeader(401)
		}

	})

}
