package main

import (
	"fmt"
	"os"
)

func main() {
	// Una empresa que se encarga de vender productos de limpieza necesita:
	// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
	// Debe tener el id del producto, precio y la cantidad.
	// Estos valores pueden ser hardcodeados o escritos en duro en una variable.

	p1 := producto{
		id:       1,
		precio:   23.41,
		cantidad: 4,
	}

	p2 := producto{
		id:       2,
		precio:   23.41,
		cantidad: 1,
	}

	var texto string
	texto = imprimir(p1, p2)
	d1 := []byte(texto)
	os.WriteFile("./productos.csv", d1, 0644)

}

type producto struct {
	id       int
	precio   float64
	cantidad int
}

func imprimir(listaProd ...producto) string {
	var cadena string
	for _, prod := range listaProd {
		cadena += fmt.Sprint(prod.id, ",", prod.precio, ",", prod.cantidad, "\n")
	}
	return cadena
}
