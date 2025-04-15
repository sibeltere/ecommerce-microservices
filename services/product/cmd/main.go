package main

import (
	"ecommerce-microservices/services/product/internal/application/services"
	"ecommerce-microservices/services/product/internal/core/storage"
	"ecommerce-microservices/services/product/internal/infrastructure/repositories"
	"ecommerce-microservices/services/product/internal/presentation/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is my first golang project")

	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=ecommerce sslmode=disable"
	storage.InitDB(dsn)

	repo := repositories.NewProductRepository(storage.DB)

	service := services.NewService(repo)

	productHandler := handlers.NewProductHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/products/{id}", productHandler.GetByID).Methods("GET")
	r.HandleFunc("/products", productHandler.GetAllProduct).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProuct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
