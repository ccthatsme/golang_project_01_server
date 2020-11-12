package main

import (
    "github.com/golang_project_01_server/services_authorization"
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

	endpoint := "authentication/authenticate"

	ds, err := datasources.NewNexientDataSource(endpoint)
	if err != nil {
		panic(err)
	}

	env := &datasources.Env{
		AuthorizationService: &services.AuthDataSource{CompanyHttp: ds},
	}

	schema := graphql.MustParseSchema(schema.GetRootSchema("./graphql/schema/schema.graphql"), &resolver.Resolver{Env: env})

	r := mux.NewRouter()
	r.HandleFunc("/auth", services_authorization.ServiceAuthorization).Methods("POST")
	r.Use(auth.Middleware())
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	http.ListenAndServe(":8081", r)

}
