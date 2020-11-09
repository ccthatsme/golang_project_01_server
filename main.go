package main

import (
	"github.com/golang_project_01_server/graphql/auth"
	"github.com/golang_project_01_server/graphql/resolver"
	"github.com/golang_project_01_server/graphql/schema"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"net/http"
)

func main() {

	schema := graphql.MustParseSchema(schema.GetRootSchema("./graphql/schema/schema.graphql"), &resolver.Resolver{})

	r := mux.NewRouter()
	r.Use(auth.Middleware())
	r.Handle("/graphql", &relay.Handler{Schema: schema})

	http.ListenAndServe(":8081", r)

}
