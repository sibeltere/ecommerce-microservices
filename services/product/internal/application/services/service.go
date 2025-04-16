package services

import (
	"ecommerce-microservices/services/product/internal/domain/entities"
	"ecommerce-microservices/services/product/internal/domain/models"
	"ecommerce-microservices/services/product/internal/infrastructure/nats"
	"fmt"
)

type ProductService struct {
	productRepository entities.Repository
}

func NewService(repo entities.Repository) *ProductService {
	return &ProductService{productRepository: repo}
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
	nats.PublishProduct(createModel)
	return result, nil
}

func (s *ProductService) UpdateProduct(updateModel *models.UpdateProductModel, Id int) (bool, error) {
	return s.productRepository.UpdateProduct(updateModel, Id)
}

func (s *ProductService) DeleteProuct(Id int) (bool, error) {
	return s.productRepository.DeleteProuct(Id)
}
