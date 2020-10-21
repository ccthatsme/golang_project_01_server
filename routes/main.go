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
Name string `json`
}

type Employee struct{

Id string `json:"id"`
FirstName string `json:"first"`
LastName string `json:"last"`

}

var newEmployee Employee

//main
func main() {

content, err := ioutil.ReadFile("employees.json")
if err != nil {
fmt.Println("could not read the file")
}

err2 := json.Unmarshal(content, &newEmployee)
if err2 != nil {
fmt.Println("unmarshal error")
}

r:=mux.NewRouter()

// newEmployee = Employee{
//             Id:"1",
//             FirstName:"Chris",
//             LastName:"Ciric"}


r.HandleFunc("/employees", getEmployees).Methods("GET")
// r.HandleFunc("/project", getProjects).Methods("GET")
// r.HandleFunc("/sow", getSow).Methods("GET")
// r.HandleFunc("/clients", getClients).Methods("GET")


http.ListenAndServe(":8081", r)

}

func getEmployees(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(newEmployee)
}





