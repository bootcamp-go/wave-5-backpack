package domain

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(id int, name string, price float64) Product {
	return Product{ID: id, Name: name, Price: price}
}
