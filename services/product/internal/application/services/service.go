package services

import (
	"ecommerce-microservices/services/product/internal/domain/entities"
	"ecommerce-microservices/services/product/internal/domain/models"
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

func (s *ProductService) CreateProduct(productModel *models.CreateProductModel) (bool, error) {
	return s.productRepository.CreateProduct(productModel)
}
