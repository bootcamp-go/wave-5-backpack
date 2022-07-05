package main

import (
	"fmt"
	"os"
)

type product struct {
	id       int
	price    float64
	quantity int
}

func newFile(productList []product) {
	var res string
	for _, product := range productList {
		res += fmt.Sprintf("%d,%.2f,%d\n", product.id, product.price, product.quantity)
	}
	os.WriteFile("products.csv", []byte(res), 0644)
}

func main() {
	insertProducts := []product{
		{
			1, 10, 100,
		},
		{
			2, 20, 10,
		},
		{
			3, 30, 1,
		},
	}
	newFile(insertProducts)
}
