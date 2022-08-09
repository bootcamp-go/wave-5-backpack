package domain

type Product struct {
	ID     int     `json:"ID"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Price  float64 `json:"price"`
	Count  int     `json:"count"`
	Code   string  `json:"code"`
	Public int8    `json:"public"`
}

type Product_Warehouse struct {
	Product
	Warehouse Warehouse `json:"warehouse"`
}
