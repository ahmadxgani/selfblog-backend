package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-blog/internal/controller"
	"go-blog/pkg/sdl"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var testSdl string

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	testSdl = `
		type Query {
			account: Account!
		}

		type Account {
			hello: String!
		}
	`
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS")))
	if err != nil {
		log.Fatalln("Error: The data source arguments are not valid")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error: Could not establish a connection with the database")
	}

}

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	// schema := graphql.MustParseSchema(sdl.Schema, &controller.Resolver{})
	schema := graphql.MustParseSchema(testSdl, &controller.Resolver{}, opts...)
	http.Handle("/query", &relay.Handler{Schema: schema})
	http.HandleFunc("/graphiql", func(w http.ResponseWriter, r *http.Request) { w.Write(sdl.Page) })
	log.Fatal(http.ListenAndServe(":8080", nil))
}
