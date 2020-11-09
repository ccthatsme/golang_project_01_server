package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/services"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
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

	var employeeType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Employee",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.ID,
				},
				"display_name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	var projectType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Project",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.ID,
				},
				"active": &graphql.Field{
					Type: graphql.Boolean,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	//schema
	fields := graphql.Fields{
		"listEmployees": &graphql.Field{
			Type:        graphql.NewList(employeeType),
			Description: "Get all employees",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return services.CheckTokenExists(services.GetEmployees), nil
			},
		},
	}

	r := mux.NewRouter()

	r.HandleFunc("/auth", services.Auth).Methods("POST")
	r.HandleFunc("/employees", services.CheckTokenExists(services.GetEmployees)).Methods("GET")
	r.HandleFunc("/employees/{employeeNetworkid}", services.CheckTokenExists(services.GetEmployee)).Methods("GET")
	r.HandleFunc("/projects", services.CheckTokenExists(services.GetProjects)).Methods("GET")
	r.HandleFunc("/projects/{id}", services.CheckTokenExists(services.GetProject)).Methods("GET")

	http.ListenAndServe(":8081", r)

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context(), "printed at handler")
	w.Header().Set("X-Authorization", services.AuthToken.Token)
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
