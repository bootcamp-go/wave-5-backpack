package main

import (
	"fmt"
	"log"
	"os"
)

type Product struct {
	id       int     `csv:"id"`
	price    float64 `csv:"price"`
	quantity int     `csv:"quantity"`
}

func main() {
	products := []Product{
		{001, 5500.30, 10},
		{011, 2500.20, 34},
		{021, 1500.10, 12},
		{031, 3500.80, 11},
		{041, 6500.50, 8},
	}

	productStore(&products)
}

func productStore(p *[]Product) {
	csvFile := fmt.Sprintln("ID, PRICE, QUANTITY")

	for _, product := range *p {
		csvFile += fmt.Sprintf("%d,%f,%d\n", product.id, product.price, product.quantity)
	}

	if err := os.WriteFile("products.csv", []byte(csvFile), 0644); err != nil {
		log.Fatal("writing file Error!")
	}

}
