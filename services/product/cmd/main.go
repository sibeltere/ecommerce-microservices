package main

import (
	"fmt"
	"log"
	"net/http"
	"productservice/internal/application/services"
	"productservice/internal/core/storage"
	natsinfra "productservice/internal/infrastructure/nats"
	"productservice/internal/infrastructure/repositories"
	"productservice/internal/presentation/handlers"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("This is my first golang project")

	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=ecommerce sslmode=disable"
	storage.InitDB(dsn)

	// setup NATS
	publisher := natsinfra.NewNatsPublisher(nats.DefaultURL)

	repo := repositories.NewProductRepository(storage.DB)

	service := services.NewService(repo, publisher)

	productHandler := handlers.NewProductHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/products/{id}", productHandler.GetByID).Methods("GET")
	r.HandleFunc("/products", productHandler.GetAllProduct).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProuct).Methods("DELETE")

	log.Println("Products service is running on port 50050")
	log.Fatal(http.ListenAndServe(":50050", r))

}
