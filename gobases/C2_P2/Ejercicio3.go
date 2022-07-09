package main

import (
	"fmt"
)

const (
	small  string = "PEQUEÑO"
	middle string = "MEDIANO"
	big    string = "GRANDE"
)

var productList []*Products

type Product interface {
	CalculateCost() float64
}

type Ecommerce interface {
	Add(price float64, productType, nameProduct string) []*Products
}

type Products struct {
	price       float64
	productType string
	name        string
}

type Shop struct {
	list []*Products
}

func newProduct(productType, nameProduct string, price float64) Product {
	return &Products{price: price, productType: productType, name: nameProduct}
	//productList = append(productList, list)
	//return list
}

func (p Products) CalculateCost() float64 {
	switch p.productType {
	case small:
		return p.price
	case middle:
		fullPrice := p.price + (p.price * 0.03)
		return fullPrice
	case big:
		fullPrice := p.price + (p.price * 0.06) + 2500
		return fullPrice
	default:
		return 0
	}
}

/*func (p Products) Add(price float64, productType, nameProduct string) *Products {
	list := &Products{
		price:       price,
		productType: productType,
		name:        nameProduct,
	}
	productList = append(productList, list)
	return list
}*/

func main() {
	p := newProduct(big, "Bicycle", 4500)
	//Ecommerce.Add(float64(2250), small, "shoes")
	fmt.Printf("El precio del producto es: %.2f\n", p.CalculateCost())
	fmt.Println(productList)
}
