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
	guardarArchivo(
		Producto{id: 223, precio: 432.22, cantidad: 2},
		Producto{id: 234, precio: 223.55, cantidad: 4},
		Producto{id: 623, precio: 1234.99, cantidad: 32},
		Producto{id: 582, precio: 4321.99, cantidad: 4},
		Producto{id: 582, precio: 2398.99, cantidad: 66},
		Producto{id: 582, precio: 7632.99, cantidad: 4},
	)

}

func guardarArchivo(productos ...Producto) error {
	var txt string
	for _, val := range productos {
		txt += fmt.Sprintf("%d, %.2f, %d;", val.id, val.precio, val.cantidad)
	}
	d1 := []byte(txt)
	err := os.WriteFile("./prod_comprados.txt", d1, 0644)
	return err
}
