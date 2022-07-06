// Ejercicio 1 - Guardar archivo
// Una empresa que se encarga de vender productos de limpieza necesita:
// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
// Debe tener el id del producto, precio y la cantidad.
// Estos valores pueden ser hardcodeados o escritos en duro en una variable.
// id,precio,cantidad
// 1,30012.00,1
// 2,1000000.00,4
// 3,50.50,1

package main

import (
	"fmt"
	"log"
	"os"
)
type Producto struct {
	Id       int
	Precio   float64
	Cantidad int
}
func suma(productos []Producto) float64 {
	var total float64
	for _, v := range productos {
		total += v.Precio * float64(v.Cantidad)
	}
	return total
}

func main() {
	productos := []Producto{
		{Id: 1, Precio: 12.00, Cantidad: 4},
		{Id: 2, Precio: 123.00, Cantidad: 5},
		{Id: 3, Precio: 132.20, Cantidad: 8},
	}

	data := "ID;Precio;Cantidad\n"
	total := fmt.Sprintf(";%.2f;;", suma(productos))
	for _, value := range productos {
		line := fmt.Sprintf("%d;%.2f;%d;\n", value.Id, value.Precio, value.Cantidad)
		data += line
	}
	data += total

	err := os.WriteFile("./productos.csv", []byte(data), 0644)
	if err != nil {
		log.Println(err)
	}
}

