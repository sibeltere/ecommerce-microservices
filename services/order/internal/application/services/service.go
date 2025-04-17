package services

import (
	"orderservice/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventPublisher interface {
	Publish(subject string, data any) (bool, error)
}

type OrderService struct {
	repository entities.Repository
	publisher  EventPublisher
}

func NewServices(repo entities.Repository, publisher EventPublisher) *OrderService {
	return &OrderService{repository: repo, publisher: publisher}
}

func (s *OrderService) GetOrder(Id primitive.ObjectID) (*entities.Order, error) {
	return s.repository.GetOrder(Id)
}
