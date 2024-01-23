package router

import (
	"github.com/marceljaworski/go-postgres_api/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products/{id:[0-9]+}", handlers.GetProduct).Methods("GET")
	router.HandleFunc("/api/products", handlers.GetAllProducts).Methods("GET")
	router.HandleFunc("/api/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id:[0-9]+}", handlers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id:[0-9]+}", handlers.DeleteProduct).Methods("DELETE")
	return router
}
