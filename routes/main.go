package main

import (
	"encoding/json"
	"fmt"
	//"github.com/friendsofgo/graphiql"
	//"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/services"
	"github.com/gorilla/mux"
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

// 	var employeeType = graphql.NewObject(
// 		graphql.ObjectConfig{
// 			Name: "Employee",
// 			Fields: graphql.Fields{
// 				"id": &graphql.Field{
// 					Type: graphql.ID,
// 				},
// 				"display_name": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 				"email": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 			},
// 		},
// 	)
//
// 	var projectType = graphql.NewObject(
// 		graphql.ObjectConfig{
// 			Name: "Project",
// 			Fields: graphql.Fields{
// 				"id": &graphql.Field{
// 					Type: graphql.ID,
// 				},
// 				"active": &graphql.Field{
// 					Type: graphql.Boolean,
// 				},
// 				"name": &graphql.Field{
// 					Type: graphql.String,
// 				},
// 			},
// 		},
// 	)
//
// 	//schema
// 	fields := graphql.Fields{
// 		"listEmployees": &graphql.Field{
// 			Type:        graphql.NewList(employeeType),
// 			Description: "Get all employees",
// 			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 				return services.CheckTokenExists(services.GetEmployees), nil
// 			},
// 		},
// 	}

// graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphiql")
// 	if err != nil {
// 		panic(err)
// 	}



	r := mux.NewRouter()

	r.Handle("/graphql", gqlHandler())
	//r.Handle("/graphiql", graphiqlHandler())



	r.HandleFunc("/auth", services.Auth).Methods("POST")
	r.HandleFunc("/employees", services.CheckTokenExists(services.GetEmployees)).Methods("GET")
	r.HandleFunc("/employees/{employeeNetworkid}", services.CheckTokenExists(services.GetEmployee)).Methods("GET")
	r.HandleFunc("/projects", services.CheckTokenExists(services.GetProjects)).Methods("GET")
	r.HandleFunc("/projects/{id}", services.CheckTokenExists(services.GetProject)).Methods("GET")
	//r.HandleFunc("/login", BasicAuth(login)).Methods("GET")
	// 	r.HandleFunc("/employees", services.CheckTokenExists(getEmployees)).Methods("GET")
	//r.HandleFunc("/employees", ValidateTokenMiddleware(getEmployees)).Methods("GET")
	//r.HandleFunc("/projects", ValidateTokenMiddleware(getProjects)).Methods("GET")
	//r.HandleFunc("/sow", ValidateTokenMiddleware(getSow)).Methods("GET")
	//r.HandleFunc("/clients", ValidateTokenMiddleware(getClients)).Methods("GET")

	http.ListenAndServe(":8081", r)

}

func gqlHandler() http.Handler{
return http.HandlerFunc(services.CheckTokenExists(services.GetEmployees))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context(), "printed at handler")
	w.Header().Set("X-Authorization", services.AuthToken.Token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmployee)
}

