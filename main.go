package main

import (
	//"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

 	"github.com/golang_project_01_server/datasources"
// 	"github.com/golang_project_01_server/repo/employee"
)

func main() {

	endpoint := "/employees"

	datasource, err := datasources.NewCompanyDatasource(endpoint)
	if err != nil {
		panic(err)
	}

	env := &datasource.Env{
		EmployeeService: &employee.EmployeeDatasource{
			CompanyHttp: datasource,
		},
	}

	schema := graphql.MustParseSchema(schema.GetRootSchema("./graphql/schema/schema.graphql"), &resolver.Resolver{Env: env})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type", "sso"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET"})

	r := mux.NewRouter()
	r.Handle("/graphql", &relay.Handler{Schema: schema})

}
