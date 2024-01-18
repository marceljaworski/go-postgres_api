package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id,omitempty`
	Message string `json:"message,omitempty`
}

// postgress values variables
const (
	host     = "localhost"
	port     = 5434
	user     = "postgres"
	dbname   = "productsdb"
	password = "postgres"
)

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro loading .env file")
	}
	password := os.Getenv("PASSWORD")

	// Postgres connection string
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := sql.Open("postgres", psqlConnString)
	if err != nil {
		panic(err)
	}

	err = db.Ping() // Check connection
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully conected to postgres")

	return db
}
