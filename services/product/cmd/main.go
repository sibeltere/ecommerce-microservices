package main

import (
	"ecommerce-microservices/services/product/internal/application/services"
	"ecommerce-microservices/services/product/internal/core/storage"
	postgrerepositories "ecommerce-microservices/services/product/internal/infrastructure/postgreRepositories"
	"ecommerce-microservices/services/product/internal/presentation/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is my first golang project")

	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=ecommerce sslmode=disable"
	storage.InitDB(dsn)

	repo := postgrerepositories.NewProductPostgresRepository(storage.DB)

	service := services.NewService(repo)

	productHandler := controllers.NewProductHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/products/{id}", productHandler.GetByID).Methods("GET")
	r.HandleFunc("/products", productHandler.GetAllProduct).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
