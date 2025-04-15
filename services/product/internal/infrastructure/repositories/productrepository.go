package repositories

import (
	"database/sql"
	"ecommerce-microservices/services/product/internal/domain/entities"
	"ecommerce-microservices/services/product/internal/domain/models"

	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetByID(id int) (*entities.Product, error) {

	var product entities.Product
	query := "SELECT id, name, quantity,price FROM product WHERE id = $1"
	row := repo.db.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.Price)
	if err != nil {
		return nil, fmt.Errorf("product not found: %v", err)
	}
	return &product, nil
}

func (repo *ProductRepository) GetAllProduct() (*[]entities.Product, error) {
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

func (repo *ProductRepository) CreateProduct(createModel *models.CreateProductModel) (bool, error) {

	query := "INSERT INTO Product (Name,Quantity,Price) Values($1,$2,$3)"

	rows, err := repo.db.Exec(query, createModel.Name, createModel.Quantity, createModel.Price)
	if err != nil {
		return false, fmt.Errorf("fail for insert %v", err)
	}
	effectedRows, er := rows.RowsAffected()

	if effectedRows < 0 || er != nil {
		return false, fmt.Errorf("fail for insert %v", err)
	}
	return true, nil
}

func (repo *ProductRepository) UpdateProduct(updateModel *models.UpdateProductModel, Id int) (bool, error) {

	query := "UPDATE Product SET Name=$1,Quantity=$2,Price=$3 WHERE Id=$4"

	rows, err := repo.db.Exec(query, updateModel.Name, updateModel.Quantity, updateModel.Price, Id)

	if err != nil {
		return false, fmt.Errorf("fail for update %v", err)
	}

	effectedRows, er := rows.RowsAffected()

	if effectedRows < 0 || er != nil {
		return false, fmt.Errorf("fail for update %v", err)
	}
	return true, nil
}

func (repo *ProductRepository) DeleteProuct(Id int) (bool, error) {

	query := "DELETE FROM Product Where Id=$1"
	rows, err := repo.db.Exec(query, Id)
	if err != nil {
		return false, fmt.Errorf("fail for delete %v", err)
	}

	effectedRows, er := rows.RowsAffected()

	if effectedRows < 0 || er != nil {
		return false, fmt.Errorf("fail for delete %v", err)
	}

	return true, nil
}
