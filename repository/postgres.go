package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/marceljaworski/go-postgres_api/lib"
	"github.com/marceljaworski/go-postgres_api/models"
)

func Insert(product models.Product) int64 {
	db, err := lib.CreateConnection()
	if err != nil {
		log.Fatalf("unable to connect the database. %v\n", err)
	}
	defer db.Close()
	sqlStatement := `INSERT INTO products(name, price, company) VALUES ($1, $2, $3) RETURNING productid`
	var id int64

	err = db.QueryRow(sqlStatement, product.Name, product.Price, product.Company).Scan(&id)

	if err != nil {
		log.Fatalf("unable to execute the query. %v\n", err)
	}

	fmt.Printf("Inserted a single record, id: %v\n", id)
	return id
}
func GetOne(id int64) (models.Product, error) {
	db, err := lib.CreateConnection()
	if err != nil {
		log.Fatalf("unable to connect the database. %v\n", err)
	}
	defer db.Close()

	var product models.Product

	sqlStatement := `SELECT * FROM products WHERE productid=$1`

	row := db.QueryRow(sqlStatement, id)

	err = row.Scan(&product.ProductID, &product.Name, &product.Price, &product.Company)

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
	db, err := lib.CreateConnection()
	if err != nil {
		log.Fatalf("unable to connect the database. %v\n", err)
	}
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
			log.Fatalf("unable to scan the row %v\n", err)
		}
		products = append(products, product)
	}
	return products, err
}

func UpdateProduct(id int64, product models.Product) int64 {
	db, err := lib.CreateConnection()
	if err != nil {
		log.Fatalf("unable to connect the database. %v\n", err)
	}
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
	db, err := lib.CreateConnection()
	if err != nil {
		log.Fatalf("unable to connect the database. %v\n", err)
	}
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
