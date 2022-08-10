package domain

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

func NewProduct(id string, name string, pType string, price float64, count int) Product {
	return Product{ID: id, Name: name, Type: pType, Price: price, Count: count}
}
