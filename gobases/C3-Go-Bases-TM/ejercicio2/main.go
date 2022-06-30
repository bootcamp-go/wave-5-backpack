package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ejercicio 2 - Leer archivo
// La misma empresa necesita leer el archivo almacenado, para ello requiere que:
// se imprima por pantalla mostrando los valores tabulados, con un título
// (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio,
// la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

func imprimirCSV(data []byte) {
	total := 0.0
	// Separamos la información leída por saltos de linea
	lineas := strings.Split(string(data), "\n")
	for _, p := range lineas {
		// Separamos cada línea por comas
		linea := strings.Split(p, ",")
		for i, l := range linea {
			switch i {
			case 0:
				// Se imprime la primera columna
				fmt.Printf("%s\t", l)
			case 1:
				// Se imprime la segunda columna
				// Se convierte el string en flotante, caso contrario, se imprime como string
				s, err := strconv.ParseFloat(l, 64)
				if err == nil {
					total += s
					fmt.Printf("%12.2f\t", s)
				} else {
					fmt.Printf("%12s\t", l)
				}
			case 2:
				// Se imprime la tercera columna
				fmt.Printf("%12s\t", l)
			}
		}
		fmt.Println("")
	}

	// Se imprime el total de los precios
	fmt.Printf("%s\t%12.2f\n\n", "TOTAL", total)
}

func main() {
	fmt.Println("Ejercicio 2 - Leer archivo")
	fmt.Println("")

	// Leemos el archivo CSV
	data, err := os.ReadFile("../ejercicio1/products.csv")

	// Verificamos que se pueda leer el archivo
	if err != nil {
		fmt.Println("Error: No se pudo leer el archivo de los productos!")
	} else {
		imprimirCSV(data)
	}
}
