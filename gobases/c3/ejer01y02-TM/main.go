package main

import (
	"fmt"
	"os"
)

// Una empresa que se encarga de vender productos de limpieza necesita:
// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
// Debe tener el id del producto, precio y la cantidad.
// Estos valores pueden ser hardcodeados o escritos en duro en una variable.

type Producto struct {
	ID       int
	Precio   float64
	Cantidad int
}

func agregarElementos(productos []Producto) {
	archivo := fmt.Sprintln("ID, Precio, Cantidad")

	for _, producto := range productos {
		archivo += fmt.Sprintf("%d, %f, %d\n", producto.ID, producto.Precio, producto.Cantidad)
	}

	err := os.WriteFile("./productos.csv", []byte(archivo), 0644)

	if err != nil {
		fmt.Println("Error")
	}
}

func main() {

	prod1 := []Producto{{1, 2.24, 1}}

	agregarElementos(prod1)

}
