package main

import (
	"log"
	"net/http"
	"orderservice/internal/application/services"
	"orderservice/internal/core/storage"
	natsinfra "orderservice/internal/infrastructure/nats"
	"orderservice/internal/infrastructure/repositories"
	"orderservice/internal/presentation/handlers"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
)

func main() {

	MONGO_URI := "mongodb://admin:secret@localhost:27017"
	storage.ConnectMongo(MONGO_URI)

	repo := repositories.NewOrderRepository(storage.Client.Database("ecommerce"))
	publisher := natsinfra.NewNatsPublisher(nats.DefaultURL)
	service := services.NewServices(repo, publisher)
	handler := handlers.NewOrderService(service)

	r := mux.NewRouter()

	r.HandleFunc("/orders/{Id}", handler.GetOrder).Methods("GET")

	log.Println("Order service is running on port 50051")
	log.Fatal(http.ListenAndServe(":50051", r))
}
