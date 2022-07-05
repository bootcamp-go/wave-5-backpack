package main

import (
	"fmt"
	"log"
	"os"
)

/*
Ejercicio 1 - Guardar archivo

Una empresa que se encarga de vender productos de limpieza necesita:
	1. Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
	2. Debe tener el id del producto, precio y la cantidad.
	3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

type Producto struct {
	Id       int
	Precio   float64
	Cantidad int
}

func main() {
	productos := []Producto{
		{Id: 111223, Precio: 30012.00, Cantidad: 1},
		{Id: 444321, Precio: 1000000.00, Cantidad: 4},
		{Id: 434321, Precio: 50.50, Cantidad: 1},
	}

	data := "Id;Precio;Cantidad\n"
	total := fmt.Sprintf(";%.2f;;", total(productos))
	for _, value := range productos {
		line := fmt.Sprintf("%d;%.2f;%d;\n", value.Id, value.Precio, value.Cantidad)
		data += line
	}
	data += total

	//fmt.Println(data)
	// linux := "xls"
	// windows := "docx"
	err := os.WriteFile("./productos.csv", []byte(data), 0644)
	if err != nil {
		log.Println(err)
	}
}

func total(productos []Producto) float64 {
	var total float64
	for _, v := range productos {
		total += v.Precio * float64(v.Cantidad)
	}
	return total
}
