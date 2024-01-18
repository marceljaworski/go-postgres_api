package router

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products/{id}", middleware.GetProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", middleware.GetAllProducts).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/products", middleware.CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/products{id}", middleware.UpdateProduct).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/products/{id}", middleware.DeleteProduct).Methods("DELETE", "OPTIONS")
}
