package main

import "fmt"

type Product interface {
	calculateCost() float64
}

type Ecommerce interface {
	Total() float64
	Add()
}

type store struct {
	list []Product
}
type product struct {
	Type  string
	Name  string
	Price float64
}

func main() {
	fmt.Println()
}
