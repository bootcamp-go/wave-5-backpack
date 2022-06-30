package main

import (
	"fmt"
	"os"
)

// Ejercicio 1 - Guardar archivo
// Una empresa que se encarga de vender productos de limpieza necesita:
//  1. Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
//     separados por punto y coma (csv).
//  2. Debe tener el id del producto, precio y la cantidad.
//  3. Estos valores pueden ser hardcodeados o escritos en duro en una variable.

// Definimos la estructura de los productos
type producto struct {
	id       int
	precio   float64
	cantidad int
}

// Función para generar la cadena de texto del CSV
func generarCSV(p []producto) string {
	// Cadena de texto para guardar la información del CVS
	prods := ""

	// Guardamos las cabeceras en la cadena de texto
	prods += fmt.Sprintf("%s,%s,%s\n", "ID", "PRECIO", "CANTIDAD")

	// Generamos la información en formato CVS para ser guardada en disco
	for _, p := range p {
		prods += fmt.Sprintf("%d,%f,%d\n", p.id, p.precio, p.cantidad)
	}

	return prods
}

// Función para guardar la cadena de texto en formato CSV en el disco
func guardarCSV(prods string) {
	// Guardamos el archivo en disco
	productosByte := []byte(prods)
	err := os.WriteFile("./products.csv", productosByte, 0644)

	// Si ocurrio un error lo mostramos al usuario
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println("Se ha guardado el archivo!")
	}
}

func main() {
	fmt.Println("Ejercicio 1 - Guardar archivo")
	fmt.Println("")

	// Definimos un arreglo de productos
	productos := []producto{
		{1, 573.0, 5},
		{2, 468210.0, 10},
		{3, 570.0, 15},
		{4, 68574.0, 20},
	}

	prods := generarCSV(productos)
	guardarCSV(prods)
}
