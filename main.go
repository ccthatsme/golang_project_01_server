package main

import (
	"github.com/golang_project_01_server/graphql/auth"
	"github.com/golang_project_01_server/graphql/resolver"
	"github.com/golang_project_01_server/graphql/schema"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/services"
)

func main() {

	endpoint := "https://portal.nexient.com/gateway/api/"

	ds, err := datasources.NewNexientDataSource(endpoint)
	if err != nil {
		panic(err)
	}

	env := &datasources.Env{
		EmployeeService: &services.EmployeeDatasource{CompanyHttp: ds},
	}

	schema := graphql.MustParseSchema(schema.GetRootSchema("./graphql/schema/schema.graphql"), &resolver.Resolver{Env: env})

	r := mux.NewRouter()
	r.Use(auth.Middleware())
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	http.ListenAndServe(":8081", r)

}
