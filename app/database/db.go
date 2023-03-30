package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	connString := fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatalln("Error: The data source arguments are not valid")
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln("Error: Could not establish a connection with the database")
	}
}
