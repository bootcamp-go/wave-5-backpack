/*
	Ejercicio 1 - Guardar archivo
	Una empresa que se encarga de vender productos de limpieza necesita: 
	- Implementar una funcionalidad para guardar un archivo de texto, con la 
	información de productos comprados, separados por punto y coma (csv).
	- Debe tener el id del producto, precio y la cantidad.
	- Estos valores pueden ser hardcodeados o escritos en duro en una variable.

	id,precio,cantidad
	1,30012.00,1
	2,1000000.00,4
	3,50.50,1
*/
package main

import (
	"fmt"
	"os"
)

func guardarArchivo(archivo []byte) {
	err := os.WriteFile("./productos.csv", archivo, 0644)
	if err != nil {
		fmt.Printf("El archivo no pudo ser creado. Error: %v", err)
	}
}

func main() {
	productosComprados := []byte("id,precio,cantidad;1, 10000.0, 2;2, 20000.0, 4;3, 30000.0, 6;")
	guardarArchivo(productosComprados)
}


