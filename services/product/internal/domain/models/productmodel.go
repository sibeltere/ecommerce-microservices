package models

type CreateProductModel struct {
	Name     string
	Quantity int
	Price    int
}

type UpdateProductModel struct {
	Name     string
	Quantity int
	Price    int
}
