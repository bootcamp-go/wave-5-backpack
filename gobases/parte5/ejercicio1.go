package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id       int
	precio   float64
	cantidad int
}

func main() {
	productos := []Producto{
		{1, 56, 2},
		{2, 51.5, 5},
		{3, 105, 0},
		{4, 213156.215, 7},
	}

	csv := []byte{}

	for _, producto := range productos {
		row := []byte(fmt.Sprintf("%d; %.2f; %d\n", producto.id, producto.precio, producto.cantidad))
		csv = append(csv, row...)
	}

	err := os.WriteFile("./productosComprados.csv", csv, 0644)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Archivo guardado correctamente.")
	}
}
