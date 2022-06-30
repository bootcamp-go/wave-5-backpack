package main

import (
	"fmt"
	"os"
)

func main() {

	p1 := newProducto(1, 100, 3)
	p2 := newProducto(2, 100, 2)
	p3 := newProducto(3, 100, 4)

	inArchivoB := []byte(fmt.Sprint(p1.detalle(), p2.detalle(), p3.detalle()))

	err := os.WriteFile("./protuctosComprados.csv", inArchivoB, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
