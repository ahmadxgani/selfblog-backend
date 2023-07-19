package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	app "go-blog/app/controller"
	"go-blog/pkg/sdl"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	connString := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln("Error: The data source arguments are not valid")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln("Error: Could not establish a connection with the database")
	}
}


func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	mux := http.NewServeMux()
	schema := graphql.MustParseSchema(sdl.Schema, &app.Resolver{}, opts...)
	mux.Handle("/query", &relay.Handler{Schema: schema})
	mux.HandleFunc("/graphiql", func(w http.ResponseWriter, r *http.Request) { w.Write(sdl.Page) })
	handler := cors.New(cors.Options{AllowedOrigins: []string{"https://studio.apollographql.com"}}).Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
