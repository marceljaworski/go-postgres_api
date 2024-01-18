package main

import (
	"fmt"
	"go-postgres_api/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	fmt.Println("Stating server on the port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
