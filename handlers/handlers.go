package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/marceljaworski/go-postgres_api/models"

	"github.com/marceljaworski/go-postgres_api/repository"

	"github.com/gorilla/mux"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatalf("Unabel to decode the request body. %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	insertID := repository.Insert(product)

	res := response{
		ID:      insertID,
		Message: "product created succesfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) // Get acces to the parameters using the mux package

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to convert the string into int. %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	product, err := repository.GetOne(int64(id))
	if err != nil {
		log.Fatalf("unable to get stock. %v\n", err)
	}
	if product.ProductID == 0 {
		w.WriteHeader(http.StatusNotFound)
		res := response{
			ID:      int64(id),
			Message: "product not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(product)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetAll()
	if err != nil {
		log.Fatalf("unable to get all the products %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(products)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("unable to convert the string into int %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	var product models.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatalf("unable to decode the request body. %v", err)
	}

	updatedRows := repository.UpdateProduct(int64(id), product)

	msg := fmt.Sprintf("Product updated successfully. Total rows/records affected %v", updatedRows)

	if updatedRows == 0 {
		msg = fmt.Sprintf("Product do not exist. Total rows/records affected %v", updatedRows)
		w.WriteHeader(http.StatusNotFound)
	}

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
		w.WriteHeader(http.StatusBadRequest)
	}

	deletedRows := repository.DeleteProduct(int64(id))

	msg := fmt.Sprintf("Product deleted suscessfully. Total rows/records %v", deletedRows)

	if deletedRows == 0 {
		msg = fmt.Sprintf("Product do not exist. Total rows/records %v", deletedRows)
		w.WriteHeader(http.StatusNotFound)
	}

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
