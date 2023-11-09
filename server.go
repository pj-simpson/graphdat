package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pj-simpson/graphdat/graph"

	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce"
	"github.com/codatio/client-sdk-go/sync-for-commerce/pkg/models/shared"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	key := os.Getenv("CODAT_API_KEY")
	if key == "" {
		panic("Please set a codat api key as an environment variable")
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CodatClient: syncforcommerce.New(
			syncforcommerce.WithSecurity(shared.Security{
				AuthHeader: key,
			}),
		),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
