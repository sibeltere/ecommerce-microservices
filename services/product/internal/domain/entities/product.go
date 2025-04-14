package entities

type Product struct {
	Id       int     `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Quantity int     `json:"quantity" db:"quantity"`
	Price    float64 `json:"price" db:"price"`
}

type Repository interface {
	GetByID(id int) (*Product, error)
	GetAllProduct() (*[]Product, error)
}
