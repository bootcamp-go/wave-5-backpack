package main

import (
	"fmt"
	"os"
)

type Productos struct {
	id       int
	precio   float64
	cantidad int
}

func pasarAString(productos []Productos) string {

	var count float32
	p := fmt.Sprintf("ID;\tPRECIO;\tCantidad\n")

	for _, producto := range productos {
		count += float32(producto.precio * float64(producto.cantidad))
		p += fmt.Sprintf("%d;\t%.2f;\t%d\n", producto.id, producto.precio, producto.cantidad)
	}

	p += fmt.Sprintf("TOTAL:\t%.2f", count)
	return p
}

func main() {

	//Ejercicio 1

	var p1 Productos = Productos{1, 3.5, 4}
	var p2 Productos = Productos{3, 7.5, 6}
	var p3 Productos = Productos{2, 8.5, 4}
	productos := []Productos{p1, p2, p3}

	paraCSV := []byte(pasarAString(productos))

	err := os.WriteFile("./productos.csv", paraCSV, 0644)

	if err != nil {
		fmt.Println("Error al escribir archivo", err)
	}

	//Ejercicio 2

	files, err := os.ReadFile("./productos.csv")

	if err != nil {
		fmt.Println("Error al leer archivo", err)
	} else {
		fmt.Println(string(files))
	}

}
