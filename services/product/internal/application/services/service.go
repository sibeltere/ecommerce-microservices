package services

import "ecommerce-microservices/services/product/internal/domain/entities"

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
