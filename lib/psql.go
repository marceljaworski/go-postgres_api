package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cenkalti/backoff"
	"github.com/joho/godotenv"
)

func CreateConnection() (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro loading .env file")
	}

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	// Postgres connection string
	psqlConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	openDB := func() error {
		db, err = sql.Open("postgres", psqlConnString)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS products (
			product_id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL,
			price INTEGER,
			company VARCHAR NOT NULL
		);`); err != nil {
		return nil, err
	}

	return db, nil
}
