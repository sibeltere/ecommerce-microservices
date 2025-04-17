package services

import (
	"ecommerce-microservices/services/product/internal/domain/entities"
	"ecommerce-microservices/services/product/internal/domain/models"
	"fmt"
)

type EventPublisher interface {
	Publish(subject string, data any) (bool, error)
}

type ProductService struct {
	productRepository entities.Repository
	publisher         EventPublisher
}

func NewService(repo entities.Repository, publisher EventPublisher) *ProductService {
	return &ProductService{productRepository: repo, publisher: publisher}
}

func (s *ProductService) GetByID(id int) (*entities.Product, error) {
	return s.productRepository.GetByID(id)
}

func (s *ProductService) GetAllProduct() (*[]entities.Product, error) {
	return s.productRepository.GetAllProduct()
}

func (s *ProductService) CreateProduct(createModel *models.CreateProductModel) (bool, error) {
	result, err := s.productRepository.CreateProduct(createModel)
	if err != nil {
		return false, fmt.Errorf("fail to create %v", err)
	}

	//publish message by nats
	s.publisher.Publish("products.created", createModel)
	return result, nil
}

func (s *ProductService) UpdateProduct(updateModel *models.UpdateProductModel, Id int) (bool, error) {
	return s.productRepository.UpdateProduct(updateModel, Id)
}

func (s *ProductService) DeleteProuct(Id int) (bool, error) {
	return s.productRepository.DeleteProuct(Id)
}
