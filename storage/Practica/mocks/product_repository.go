package mock

import "github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"

var MockProduct domain.Product = domain.Product{
	Name:         "test",
	Color:        "Red",
	Price:        10.99,
	Stock:        10,
	Code:         "JH7BU998G",
	Published:    true,
	Warehouse_id: 1,
}

var MockProductList []domain.Product = []domain.Product{
	{
		Id:           1,
		Name:         "product 1",
		Color:        "red",
		Price:        10.99,
		Stock:        100,
		Code:         "HJ988BH",
		Created_at:   "2022-08-10",
		Published:    true,
		Warehouse_id: 1,
	},
}

var MockProductOne domain.Product = domain.Product{
	Id:           1,
	Name:         "product 1",
	Color:        "red",
	Price:        10.99,
	Stock:        100,
	Code:         "HJ988BH",
	Created_at:   "2022-08-10",
	Published:    true,
	Warehouse_id: 1,
}

var MockProductEmpty domain.Product = domain.Product{}
