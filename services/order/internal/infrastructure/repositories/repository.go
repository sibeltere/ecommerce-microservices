package repositories

import (
	"context"
	"orderservice/internal/domain/entities"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{collection: db.Collection("orders")}
}

func (r *OrderRepository) GetOrder(Id primitive.ObjectID) (*entities.Order, error) {
	var order entities.Order
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": Id}).Decode(&order)
	if err != nil {
		return nil, fmt.Errorf("order not found: %v", err)
	}

	return &order, nil
}
