package main

import (
	"log"
	"net/http"

	app "go-blog/app/controller"
	"go-blog/pkg/sdl"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/lib/pq"
)

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(sdl.Schema, &app.Resolver{}, opts...)
	http.Handle("/query", &relay.Handler{Schema: schema})
	http.HandleFunc("/graphiql", func(w http.ResponseWriter, r *http.Request) { w.Write(sdl.Page) })
	log.Fatal(http.ListenAndServe(":8080", nil))
}
