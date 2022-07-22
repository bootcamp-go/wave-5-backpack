package main

import (
	"fmt"
	"os"
)

type Product struct {
	id       int
	price    float32
	quantity int
}

func storedFile(file string, products []Product) bool {
	data := "ID;PRICE;QUANTITY\n"
	for _, prod := range products {
		data = data + fmt.Sprintf("%d;%.2f;%d\n", prod.id, prod.price, prod.quantity)
	}
	err := os.WriteFile(file, []byte(data), 0644)
	if err != nil {
		return false
	} else {
		return true
	}
}

func main() {
	products := []Product{
		{id: 1, price: 65.2, quantity: 5},
		{id: 2, price: 118, quantity: 10},
		{id: 3, price: 430.5, quantity: 25},
	}
	file := "./products.csv"
	err := storedFile(file, products)
	if !err {
		fmt.Println("El Archivo no se pudo escribir")
	} else {
		fmt.Println("Archivo escrito con Productos")
	}

}
