package router

import (
	"go-postgres_api/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products/{id}", handlers.GetProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", handlers.GetAllProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", handlers.CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/products{id}", handlers.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/products/{id}", handlers.DeleteProduct).Methods("DELETE", "OPTIONS")
}
