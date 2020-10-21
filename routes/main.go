package main

import (
"fmt"
"github.com/gorilla/mux"
"encoding/json"
"net/http"
"io/ioutil"

)

type Project struct{

Id string `json:"id"`
ProjectName string `json:"projectName"`

}

type Employee struct{

Id string `json:"id"`
FirstName string `json:"first"`
LastName string `json:"last"`

}

var newEmployee Employee
var newProject Project

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

r:=mux.NewRouter()


r.HandleFunc("/employees", getEmployees).Methods("GET")
r.HandleFunc("/projects", getProjects).Methods("GET")
// r.HandleFunc("/sow", getSow).Methods("GET")
// r.HandleFunc("/clients", getClients).Methods("GET")


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





