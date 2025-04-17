package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Total float64            `bson:"total"`
	Items []*OrderItem       `bson:"items"`
}

type OrderItem struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	ProductId int                `bson:"productId"`
	Quantity  int                `bson:"quantity"`
	Price     float64            `bson:"price"`
}

type Repository interface {
	GetOrder(Id primitive.ObjectID) (*Order, error)
	// GetOrders() (*[]Order, error)
	// CreateOrder(createModel *models.CreateOrderModel) (bool, error)
	// UpdateOrder(updateModel *models.UpdateOrderModel) (bool, error)
	// DeleteOrder(Id int) (bool, error)
}
