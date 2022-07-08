package main

import (
	"fmt"
)

/*Ejercicio 1 - Letras de una palabra

La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada
una de las letras por separado para deletrearla.
	1. Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
	2. Luego imprimí cada una de las letras.*/

func main() {

	var palabra string = "Bootcamp"
	// fmt.Println(palabra[1]) // devuelve una runa - el tipo char de Go
	// Por esta razón se realiza el cast to string

	fmt.Printf("Número de letras: %d\n", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%s ", string(palabra[i]))
	}
	fmt.Println()

	// Otra forma
	for _, value := range palabra {
		fmt.Printf("%s ", string(value))
	}
	fmt.Println()

}
