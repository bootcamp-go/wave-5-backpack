package main

import "fmt"

const (
	big    = "big"
	medium = "medium"
	small  = "small"
)

type productInter interface {
	calcCost() float64
}

type ecommerce interface {
	total() float64
	add(p productInter)
}

type product struct {
	typ   string
	name  string
	price float64
}

func (p product) calcCost() float64 {
	switch p.typ {
	case big:
		return p.price*1.06 + 2500
	case medium:
		return p.price * 1.03
	case small:
		return p.price
	}
	return 0
}

func newProduct(typ string, name string, price float64) *product {
	return &product{typ, name, price}
}

type shop struct {
	products []productInter
}

func (s shop) total() float64 {
	result := 0.0
	for _, value := range s.products {
		result += value.calcCost()
	}
	return result
}

func (s *shop) add(p productInter) {
	s.products = append(s.products, p)
}

func newShop() ecommerce {
	return &shop{}
}

func main() {
	e := newShop()
	e.add(newProduct(big, "tt1", 300))
	e.add(newProduct(medium, "tt1", 300))
	e.add(product{small, "tt1", 300})

	fmt.Println(e.total())
}
