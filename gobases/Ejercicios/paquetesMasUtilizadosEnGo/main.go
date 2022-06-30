package main

import (
	"fmt"
	"os"
)

func main() {
	/*//1
	p1 := newProducto(1, 100, 3)
	p2 := newProducto(2, 100, 2)
	p3 := newProducto(3, 100, 4)

	inArchivoB := []byte(fmt.Sprint(p1.detalle(), p2.detalle(), p3.detalle()))

	err := os.WriteFile("./productosComprados.csv", inArchivoB, 0644)

	if err != nil {
		fmt.Println(err)
	}*/

	//2
	archivo, err := os.ReadFile("productosComprados.csv")

	if err != nil {
		fmt.Println(err)
	} else {
		mostrarEnPantalla(archivo)
	}
}
