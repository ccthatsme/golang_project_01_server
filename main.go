package main

import (
// 	"encoding/json"
// 	"fmt"
	//"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/resolver"
	"github.com/golang_project_01_server/graphql/schema"
	"github.com/golang_project_01_server/services"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	//"github.com/graphql-go/graphql"
	//"io/ioutil"
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

	schema := graphql.MustParseSchema(schema.GetRootSchema("./graphql/schema/schema.graphql"), &resolver.Resolver{})
// 	allowedHeaders := handlers.AllowedHeaders([]string{"content-type", "sso"})
// 	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
// 	allowedMethods := handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET"})

	r := mux.NewRouter()
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	r.HandleFunc("/auth", services.Auth).Methods("POST")
	r.HandleFunc("/employees", services.CheckTokenExists(services.GetEmployees)).Methods("GET")
	r.HandleFunc("/employees/{employeeNetworkid}", services.CheckTokenExists(services.GetEmployee)).Methods("GET")
	r.HandleFunc("/projects", services.CheckTokenExists(services.GetProjects)).Methods("GET")
	r.HandleFunc("/projects/{id}", services.CheckTokenExists(services.GetProject)).Methods("GET")

	http.ListenAndServe(":8081", r)

}
