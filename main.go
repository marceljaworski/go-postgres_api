package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/marceljaworski/go-postgres_api/lib"
	"github.com/marceljaworski/go-postgres_api/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro loading .env file")
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r := router.Router()
	r.Use(lib.LoggingMiddleware)

	fmt.Printf("Starting server on the port %v\n", httpPort)

	log.Fatal(http.ListenAndServe(":"+httpPort, r))
}
