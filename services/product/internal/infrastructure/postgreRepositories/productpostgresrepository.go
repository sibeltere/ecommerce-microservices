package postgrerepositories

import (
	"database/sql"
	"ecommerce-microservices/services/product/internal/domain/entities"

	"fmt"
)

type ProductPostgresRepository struct {
	db *sql.DB
}

func NewProductPostgresRepository(db *sql.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{db: db}
}

func (repo *ProductPostgresRepository) GetByID(id int) (*entities.Product, error) {

	var product entities.Product
	query := "SELECT id, name, quantity,price FROM product WHERE id = $1"
	row := repo.db.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("product not found: %v", err)
	}
	return &product, nil
}

func (repo *ProductPostgresRepository) GetAllProduct() (*[]entities.Product, error) {
	var productList []entities.Product
	query := "Select * from product"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("there is no any product %v", err)
	}

	for rows.Next() {
		var p entities.Product
		err := rows.Scan(&p.Id, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, fmt.Errorf("there is no any product %v", err)
		}
		productList = append(productList, p)
	}
	return &productList, nil
}
