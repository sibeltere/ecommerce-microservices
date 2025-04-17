package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectMongo(uri string) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Mongo connection error %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("ping was not attach for Mongo %v", err)
	}

	log.Printf("Mongo connection is succesfull")
	Client = client
}
