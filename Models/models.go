package Models

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

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
}

var Employees []Employee

//change to project at some point
type Client struct {
	Id             string `json:"id"`
	ClientName     string `json:"clientName"`
	ContractLength int    `json:"contractLength"`
}