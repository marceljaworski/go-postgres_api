package repository

import (
	"database/sql"
	"fmt"
	"go-postgres_api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// postgress values variables
const (
	host   = "localhost"
	port   = 5434
	user   = "postgres"
	dbname = "productsdb"
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

func Insert(product models.Product) int64 {
	db := createConnection()
	defer db.Close()
	sqlStatement := `INSERT INTO products(name, price, company) VALUES ($1, $2, $3) RETURNING productid`
	var id int64

	err := db.QueryRow(sqlStatement, product.Name, product.Price, product.Company).Scan(&id)

	if err != nil {
		log.Fatalf("unable to execute the query. %v\n", err)
	}

	fmt.Printf("Inseted a single record, id:%v\n", id)
	return id
}
func GetOne(id int64) (models.Product, error) {
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
		log.Fatalf("unable to scan the ro. %v\n", err)
	}

	return product, err
}
func GetAll() ([]models.Product, error) {
	db := createConnection()
	defer db.Close()

	var products []models.Product

	sqlStatement := `SELECT * FROM products`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("unable to execute the query. %v\n", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ProductID, &product.Name, &product.Price, &product.Company)
		if err != nil {
			log.Fatalf("ufnable to scan the row %v\n", err)
		}
		products = append(products, product)
	}
	return products, err
}

func UpdateProduct(id int64, product models.Product) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `UPDATE products SET name=$2, price=$3, company=$4 WHERE productid=$1`

	res, err := db.Exec(sqlStatement, id, product.Name, product.Price, product.Company)

	if err != nil {
		log.Fatalf("unable to execute the query %v\n", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("error whilw checking the affected rows. %v\n", err)
	}

	fmt.Printf("Total rows/records affected %v\n", rowsAffected)

	return rowsAffected
}

func DeleteProduct(id int64) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM products WHERE productid=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("unable to execute the query %v\n", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("error whilw checking the affected rows. %v\n", err)
	}

	fmt.Printf("Total rows/records affected %v\n", rowsAffected)

	return rowsAffected
}
