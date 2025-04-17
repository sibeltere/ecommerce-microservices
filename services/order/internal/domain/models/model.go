package models

type CreateOrderModel struct {
	Total float64
	Items []*Item
}

type UpdateOrderModel struct {
	Total float64
	Items []*Item
}

type Item struct {
	ProductId int
	Quantity  int
	Price     float64
}
