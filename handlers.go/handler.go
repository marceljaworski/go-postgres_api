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

func createConnection() *sql.DB {
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
	products, err := getAllProducts()
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
	id, err := strconv.Atoi(params["id"])
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

func insertProduct(product models.Product) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO products(name, price, company) VALUES ($1, $2, $3) RETURNING productid`
	var id int64

	err := db.QueryRow(sqlStatement, product.Name, product.Price, product.Company).Scan(&id)

	if err != nil {
		log.Fatal("unable to execute the query. %v", err)
	}

	fmt.Printf("Inseted a single record %v", id)
	return id
}

func getProduct(id int64) (models.Product, error) {
	db := createConnection()
	defer db.Close()

	var product models.Product

	sqlStatement := `SELECT * FROM products WHERE productid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&product.ProductID, &product.Name, &product.Price, &product.Company)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return product, nil
	case nil:
		return product, nil
	default:
		log.Fatalf("unable to scan the ro. %v", err)
	}

	return product, err
}

func getAllProducts() ([]models.Product, error) {
	db := createConnection()
	defer db.Close()

	var products []models.Product

	sqlStatement := `SELECT * FROM products`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("unable to execute the query. %v, err")
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ProductID, &product.Name, &product.Price, &product.Company)
		if err != nil {
			log.Fatalf("ufnable to scan the row %v", err)
		}
		products = append(products, product)
	}
	return products, err
}

func updateProduct(id int64, product models.Product) int64 {

}

func deleteProduct(id int64) int64 {

}
