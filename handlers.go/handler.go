package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go-postgres_api/models"

	"github.com/gorilla/mux"
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

	fmt.Println("Successfully connected to postgres")

	return db
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatal("Unabel to decode the request body. %v", err)
	}

	insertID := insertProduct(product)

	res := response{
		ID:      insertID,
		Message: "product created succesfully",
	}

	json.NewEncoder(w).Encode(res)
}
func GetProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) // Get acces to the parameters using the mux package

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to convert the string into int. %v", err)
	}

	product, err := getProduct(int64(id))
	if err != nil {
		log.Fatalf("unable to get stock. %v", err)
	}

	json.NewEncoder(w).Encode(product)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := GetAllProducts()
	if err != nil {
		log.Fatalf("unable to get all the products %v", err)
	}

	json.NewEncoder(w).Encode(products)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to convert the string into int %v", err)
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatalf("unable to decode the request body. %v", err)
	}

	updatedRows := updateProduct(int64(id), product)

	msg := fmt.Sprintf("Product updated successfully. Total rows/records affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"])
	if err != nil {
		log.Fatalf("unable to convert string to int. %v", err)
	}

	deletedRows := deleteProduct(int64(id))

	msg := fmt.Sprintf("Product deleted suscessfully. Total rows/records %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
