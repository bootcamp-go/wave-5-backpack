package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id       int     `csv:"id"`
	precio   float64 `csv:"precio"`
	cantidad int     `csv:"cantidad"`
}

func guardarProductos(productos []Producto) {
	csvData := fmt.Sprintln("id,precio,cantidad")

	for _, producto := range productos {
		csvData += fmt.Sprintf("%d,%f,%d\n", producto.id, producto.precio, producto.cantidad)
	}

	err := os.WriteFile("productos.csv", []byte(csvData), 0644)

	if err != nil {
		fmt.Println("Error al escribir el archivo")
	}
}

func main() {
	productos := []Producto{
		{100, 115000.0, 10},
		{200, 120000.0, 20},
		{300, 130000.0, 30},
		{400, 140000.0, 40},
	}

	guardarProductos(productos)

}
