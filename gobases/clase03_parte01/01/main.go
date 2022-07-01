/*
Ejercicio 1 - Guardar archivo
Una empresa que se encarga de vender productos de limpieza necesita: 
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.


*/

package main

import (
	"fmt"
	//"io"
	"os"
)

type Productos struct {
	id int
	precio float64
	cantidad int
}

func pasarAString(productos []Productos) string {
	var total float64
	datos := fmt.Sprintf("ID;\tPRECIO;\tCANT\n")

	for _, p := range productos {
		total += (p.precio*float64(p.cantidad))
		datos += fmt.Sprintf("%d;\t%.2f;\t%d\n",p.id,p.precio,p.cantidad)
	}

	datos += fmt.Sprintf("TOTAL:\t%.2f", total)

	return datos
}


func main(){

	p1 := Productos{id: 1, precio:24.0, cantidad:5}
	p2 := Productos{id: 2, precio:33.4, cantidad:8}
	p3 := Productos{id: 3, precio:12.5, cantidad:2}

	productos := []Productos{p1,p2,p3}

	data := []byte(pasarAString((productos)))

	err := os.WriteFile("./productos.csv", data, 0644)

	if err != nil {
		return
	}

}

