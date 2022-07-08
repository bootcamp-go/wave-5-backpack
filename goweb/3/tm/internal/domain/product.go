package domain

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func NewProduct(id int, name string, price float64, quantity int) Product {
	return Product{ID: id, Name: name, Price: price, Quantity: quantity}
}
