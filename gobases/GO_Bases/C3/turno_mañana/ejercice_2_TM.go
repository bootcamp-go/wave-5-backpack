package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./products.csv")
	if err != nil {
		fmt.Println("Ocurrio un error al leer el archivo products.csv")
		os.Exit(1)
	}
	products := string(data)
	fmt.Println(products)
}
