package main

import (
	"fmt"
	"os"
)

func main() {
	b := []byte("Id Producto;Precio;Cantidad \n   1;2000;1 \n   2;4000;4 \n   4;1000;3")
	err := os.WriteFile("Productos.csv", b, 0644)
	if err != nil {
		fmt.Printf("Error escritura: %v", err)
	}

	// Ejercicio 2 Leer archivo

	data, err := os.ReadFile("./Productos.csv")

	if err != nil {
		fmt.Printf("Error lectura: %v", err)
	}
	fmt.Printf("file: %v \n \n", string(data))

}
