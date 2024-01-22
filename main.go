package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marceljaworski/go-postgres_api/router"
)

func main() {
	r := router.Router()

	fmt.Println("Stating server on the port 8081")

	log.Fatal(http.ListenAndServe(":8081", r))
}
