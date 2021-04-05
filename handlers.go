package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
	"sa-back/graph"
	"sa-back/graph/generated"
)

func gqlHandler() *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
}

func handlers(mux *http.ServeMux) {
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", gqlHandler())
	mux.HandleFunc("/upload", uploadHandler)
}
