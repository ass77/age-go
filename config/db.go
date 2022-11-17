package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/ass77/age-go/age"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	var graphName string = "working_person"

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var dsn = os.Getenv("DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	// GetReady prepare AGE extension load AGE extension set graph path
	_, err = age.GetReady(db, graphName)
	if err != nil {
		panic(err)
	}

	log.Println("AGE extension loaded")

	return db

}
