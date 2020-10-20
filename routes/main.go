package main

import (
"fmt"
"github.com/gorilla/mux"
"encoding/json"
"net/http"

)

//structs
type Employee struct{

Id string `json:id`
FirstName string `json:first`
LastName string `json:first`

}

var newEmployee Employee

var employees []Employee

//main
func main() {

r:=mux.NewRouter()

newEmployee = Employee{
            Id:"1",
            FirstName:"Chris",
            LastName:"Ciric"}

employees = append(employees, newEmployee)
fmt.Println(newEmployee)

r.HandleFunc("/employees", getEmployees).Methods("GET")
// r.HandleFunc("/project", getProjects).Methods("GET")
// r.HandleFunc("/sow", getSow).Methods("GET")
// r.HandleFunc("/clients", getClients).Methods("GET")

http.ListenAndServe(":8081", r)

}

func getEmployees(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
fmt.Println("new test")
fmt.Println(newEmployee)
json.NewEncoder(w).Encode(newEmployee)
}




