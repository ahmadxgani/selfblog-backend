package main

import (
	"log"
	"net/http"

	app "go-blog/app/controller"
	"go-blog/pkg/sdl"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	mux := http.NewServeMux()
	schema := graphql.MustParseSchema(sdl.Schema, &app.Resolver{}, opts...)
	mux.Handle("/query", &relay.Handler{Schema: schema})
	mux.HandleFunc("/graphiql", func(w http.ResponseWriter, r *http.Request) { w.Write(sdl.Page) })
	handler := cors.New(cors.Options{AllowedOrigins: []string{"https://studio.apollographql.com"}}).Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
