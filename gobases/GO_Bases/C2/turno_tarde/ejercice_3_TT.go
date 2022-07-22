package main

const (
	small  = "SMALL"
	medium = "MEDIUM"
	big    = "BIG"
)

type Product struct {
	product_type string
	name         string
	price        float32
}

type Store struct {
	products []Product
}

type productInterface interface {
	calculated_cost() float32
}

type ecommerceInterface interface {
	addProduct(product Product)
	total() float32
}

func (p Product) calculated_cost() float32 {
	switch p.product_type {
	case "SMALL":
		return 0
	case "MEDIUM":
		return p.price * 0.03
	case "BIG":
		return p.price*0.03 + 2500
	default:
		return 0
	}
}

func (s *Store) addProduct(product Product) {
	s.products = append(s.products, product)
}

func (s *Store) total() float32 {
	var sum float32 = 0
	for _, p := range s.products {
		sum += p.price + p.calculated_cost()
	}
	return sum
}

func newProduct(product_type, name string, price float32) productInterface {
	p := Product{product_type, name, price}
	return p
}

func newStore() ecommerceInterface {
	s := Store{products: []Product{}}
	return &s
}

func main() {

}
