package main

import (
"fmt"
"github.com/gorilla/mux"
"encoding/json"
"net/http"
"io/ioutil"

)

type Client struct {

Id string `json:"id"`
ClientName string `json:"clientName"`
ContractLength int `json:"contractLength"`

}

type Project struct{

Id string `json:"id"`
ProjectName string `json:"projectName"`

}

type SOW struct{

Id string `json:"id"`
WorkName string `json:"workName"`

}

type Employee struct{

Id string `json:"id"`
FirstName string `json:"first"`
LastName string `json:"last"`

}

var newEmployee Employee
var newProject Project
var newSow SOW
var newClient Client

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

r:=mux.NewRouter()


r.HandleFunc("/employees", getEmployees).Methods("GET")
r.HandleFunc("/projects", getProjects).Methods("GET")
r.HandleFunc("/sow", getSow).Methods("GET")
r.HandleFunc("/clients", getClients).Methods("GET")


http.ListenAndServe(":8081", r)

}

func getEmployees(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(newEmployee)
}

func getProjects(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(newProject)
}

func getSow(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(newSow)
}

func getClients(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(newClient)
}




