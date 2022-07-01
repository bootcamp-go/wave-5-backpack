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
		{1, 60.0, 1},
		{2, 10.0, 1},
		{3, 5.5, 1},
		{3, 5000.0, 2},
	}

	guardarArchivo(productos)
}

func guardarArchivo(productos []Producto) {
	var productosConcatenados string

	for _, producto := range productos {
		productosConcatenados = productosConcatenados + productoToString(producto)
	}
	registrosBytes := []byte(productosConcatenados)
	os.WriteFile("./myFile.txt", registrosBytes, 0644)
}

func productoToString(producto Producto) string {
	registro := fmt.Sprintln(producto.id, ",", producto.precio, ",", producto.cantidad)
	return registro
}
