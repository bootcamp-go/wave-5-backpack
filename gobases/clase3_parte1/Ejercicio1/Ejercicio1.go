package main

import (
	"fmt"
	"os"
)

type Producto struct {
	ID       int
	Precio   int
	Cantidad int
}

func main() {
	producto1 := Producto{
		ID:       1,
		Precio:   10,
		Cantidad: 5,
	}
	producto2 := Producto{
		ID:       2,
		Precio:   8,
		Cantidad: 4,
	}
	producto3 := Producto{
		ID:       3,
		Precio:   35,
		Cantidad: 2,
	}
	fmt.Println(imprimirConSeparador(',', "../result.csv", producto1, producto2, producto3))
}

func imprimirConSeparador(separador rune, path string, productos ...Producto) error {
	var strProducto string
	for _, producto := range productos {
		strProducto += fmt.Sprintf("%d,%d,%d\n", producto.ID, producto.Precio, producto.Cantidad)
	}
	byteArray := []byte(strProducto)
	err := os.WriteFile(path, byteArray, 0644)
	return err
}
