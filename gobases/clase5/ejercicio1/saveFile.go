package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id       int
	precio   float64
	cantidad int
	total    float64
}

func main() {
	prod1 := Producto{1, 23, 4, 0}
	prod2 := Producto{2, 23, 5, 0}

	dataProductos := []byte(fmt.Sprintf("%v \n %v", prod1, prod2))

	err := os.WriteFile("./productos.csv", dataProductos, 0644)

	if err != nil {
		fmt.Printf("Ocurrio un error: %v \n", err)
	}
}
