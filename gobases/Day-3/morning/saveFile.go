package main

import (
	"fmt"
	"os"
)

type Product struct {
	id     int
	name   string
	price  float64
	amount int
}

func writeCsvFile(products []Product, fileName string) {
	finalCsv := "Id,Name,Price,Amount\n"
	for _, product := range products {
		finalCsv += fmt.Sprintf("%d,%s,%2.f,%d\n", product.id, product.name, product.price, product.amount)
	}
	csvToByte := []byte(finalCsv)
	os.WriteFile(fmt.Sprintf("./%s.csv", fileName), csvToByte, 0644)
}

func main() {
	products := []Product{
		{1, "Cacao", 2.0, 1},
		{1, "Cheese", 4.0, 1},
	}
	writeCsvFile(products, "products")
}
