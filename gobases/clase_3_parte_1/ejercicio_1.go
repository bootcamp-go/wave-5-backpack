package main

import (
	"fmt"
	"os"
)

type Producto struct {
	id       int
	precio   int
	cantidad int
}

func escribirProductosACsv(productos []Producto, separador string) {
	datos := "id;precio;cantidad \n"
	for _, producto := range productos {
		datos += fmt.Sprint(producto.id, separador, producto.precio, separador, producto.cantidad, "\n")
	}
	os.WriteFile("./productos.csv", []byte(datos), 0644)
}

func main() {
	productos := []Producto{{
		id:       1,
		precio:   200,
		cantidad: 4,
	}, {
		id:       2,
		precio:   300,
		cantidad: 8,
	}, {
		id:       3,
		precio:   500,
		cantidad: 9,
	}}
	escribirProductosACsv(productos, ";")
	os.Exit(1)
}
