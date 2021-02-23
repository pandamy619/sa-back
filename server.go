package main

import (
	"fmt"
	"log"
	"net/http"
	"sa-back/graph"
	"sa-back/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	confFile     = "config"
	pathConfFile = "."
	typeConfFile = "yml"
)

func main() {
	config, err := loadEnv(confFile, pathConfFile, typeConfFile)
	if err != nil {
		fmt.Println(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
